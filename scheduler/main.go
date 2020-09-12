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
	middleware.InitProducer(ctx, os.GetEnv("BROKERS"))
	models.InitDB(os.GetEnv("DB_PASSWD"), os.GetEnv("DB_HOST"))
	crontask.Start(ctx, 100, time.Minute*time.Duration(10), scanner.NewScanner())
}
