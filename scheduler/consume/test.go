package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

func main() {
	sarama.Logger = log.New(os.Stdout, "[sarama] ", log.LstdFlags)

	config := sarama.NewConfig()
	//接收失败通知
	config.Consumer.Return.Errors = true
	//设置使用的kafka版本,如果低于V0_10_0_0版本,消息中的timestrap没有作用.需要消费和生产同时配置
	config.Version = sarama.V1_1_0_0
	//新建一个消费者
	consumer, e := sarama.NewConsumer([]string{"172.30.39.100:9092"}, config)
	if e != nil {
		panic("error get consumer")
	}
	defer consumer.Close()

	//根据消费者获取指定的主题分区的消费者,Offset这里指定为获取最新的消息.
	partitionConsumer, err := consumer.ConsumePartition("import_raw", 0, sarama.OffsetNewest)
	if err != nil {
		fmt.Println("error get partition consumer", err)
	}
	defer partitionConsumer.Close()
	//循环等待接受消息.
	for {
		select {
		//接收消息通道和错误通道的内容.
		case msg := <-partitionConsumer.Messages():
			fmt.Println("msg offset: ", msg.Offset, " partition: ", msg.Partition, " timestrap: ", msg.Timestamp.Format("2006-Jan-02 15:04"), " value: ", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			fmt.Println(err.Err)
		}
	}
}
