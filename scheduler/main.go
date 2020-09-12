package main

import (
	"context"
	"time"

	"github.com/piggona/fundings_view/scheduler/crontask"
	"github.com/piggona/fundings_view/scheduler/middleware"
	"github.com/piggona/fundings_view/scheduler/scanner"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	brokers := []string{"172.30.39.100:9092"}
	middleware.InitProducer(ctx, brokers)
	crontask.Start(ctx, 100, time.Minute*time.Duration(10), scanner.NewScanner())
}
