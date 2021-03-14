package crontask

import (
	"context"
	"time"
)

type Worker struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration, r *Runner) *Worker {
	return &Worker{
		ticker: time.NewTicker(interval),
		runner: r,
	}
}

func (w *Worker) startWorker(ctx context.Context) {
	for {
		select {
		case <-w.ticker.C:
			go w.runner.StartAll()
		case <-ctx.Done():
			return
		}
	}
}

func Start(ctx context.Context, bufSize int, interval time.Duration, exec ExecPair) {
	r := NewRunner(bufSize, exec)
	worker := NewWorker(interval, r)
	worker.startWorker(ctx)
}
