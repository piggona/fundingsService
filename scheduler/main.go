package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/piggona/fundingsService/scheduler/crontask"
	"github.com/piggona/fundingsService/scheduler/middleware"
	"github.com/piggona/fundingsService/scheduler/models"
	"github.com/piggona/fundingsService/scheduler/scanner"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	middleware.InitProducer(ctx, splitEnv(os.Getenv("BROKERS")))
	models.InitDB(os.Getenv("DB_PASSWD"), os.Getenv("DB_HOST"))
	crontask.Start(ctx, 100, time.Minute*time.Duration(10), scanner.NewScanner())
}

func splitEnv(str string) []string {
	return strings.Split(str, ",")
}
