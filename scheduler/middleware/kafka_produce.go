package middleware

import (
	"context"
	logger "log"
	"os"
	"strings"

	"github.com/Shopify/sarama"
	"github.com/piggona/fundingsService/scheduler/utils/log"
	"go.opentelemetry.io/otel/api/global"
	"go.opentelemetry.io/otel/api/kv"
	"go.opentelemetry.io/otel/api/trace"
)

const (
	MessageIDHeaderName = "message_id"
	SpanHeaderName      = "span_id"
	TraceHeaderName     = "trace_id"
)

type OTelInterceptor struct {
	tracer     trace.Tracer
	fixedAttrs []kv.KeyValue
}

func (oi *OTelInterceptor) OnSend(msg *sarama.ProducerMessage) {
	if shouldIgnoreMsg(msg) {
		return
	}
	_ = oi.tracer.WithSpan(context.TODO(), msg.Topic,
		func(ctx context.Context) error {
			span := trace.SpanFromContext(ctx)
			spanContext := span.SpanContext()
			attWithTopic := append(
				oi.fixedAttrs,
				kv.String("messaging.destination", msg.Topic),
				kv.String("messaging.message_id", spanContext.SpanID.String()),
			)
			span.SetAttributes(attWithTopic...)

			// remove existing partial tracing headers if exists
			noTraceHeaders := msg.Headers[:0]
			for _, h := range msg.Headers {
				key := string(h.Key)
				if key != TraceHeaderName && key != SpanHeaderName && key != MessageIDHeaderName {
					noTraceHeaders = append(noTraceHeaders, h)
				}
			}
			traceHeaders := []sarama.RecordHeader{
				{Key: []byte(TraceHeaderName), Value: []byte(spanContext.TraceID.String())},
				{Key: []byte(SpanHeaderName), Value: []byte(spanContext.SpanID.String())},
				{Key: []byte(MessageIDHeaderName), Value: []byte(spanContext.SpanID.String())},
			}
			msg.Headers = append(noTraceHeaders, traceHeaders...)
			return nil
		})
}

func shouldIgnoreMsg(msg *sarama.ProducerMessage) bool {
	// check message hasn't been here before (retries)
	var traceFound, spanFound, msgIDFound bool
	for _, h := range msg.Headers {
		if string(h.Key) == TraceHeaderName {
			traceFound = true
			continue
		}
		if string(h.Key) == SpanHeaderName {
			spanFound = true
			continue
		}
		if string(h.Key) == MessageIDHeaderName {
			msgIDFound = true
		}
	}
	return traceFound && spanFound && msgIDFound
}

func NewOTelInterceptor(brokers []string) *OTelInterceptor {
	oi := OTelInterceptor{}
	oi.tracer = global.TraceProvider().Tracer("shopify.com/sarama/examples/interceptors")

	// These are based on the spec, which was reachable as of 2020-05-15
	// https://github.com/open-telemetry/opentelemetry-specification/blob/master/specification/trace/semantic_conventions/messaging.md
	oi.fixedAttrs = []kv.KeyValue{
		kv.String("messaging.destination_kind", "topic"),
		kv.String("span.otel.kind", "PRODUCER"),
		kv.String("messaging.system", "kafka"),
		kv.String("net.transport", "IP.TCP"),
		kv.String("messaging.url", strings.Join(brokers, ",")),
	}
	return &oi
}

var producer sarama.SyncProducer

func InitProducer(ctx context.Context, brokers []string) {
	var err error
	sarama.Logger = logger.New(os.Stdout, "[Sarama]", logger.LstdFlags)
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll // 发送完数据需要leader和follow都确认
	// config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true // 成功交付的消息将在success channel返回

	producer, err = sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Error("producer closed, err:", err)
		return
	}
	go func() {
		<-ctx.Done()
		producer.Close()
		return
	}()
}

func PutData(topic string, message string) {
	log.Info("[Sarama] put data into topic %s, message: %s", topic, message)
	msg := &sarama.ProducerMessage{}
	msg.Topic = topic
	msg.Value = sarama.StringEncoder(message)
	pid, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Error("send msg failed, err:", err)
		return
	}
	log.Info("pid:%v offset:%v\n", pid, offset)
}
