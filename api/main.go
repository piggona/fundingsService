package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	srv := createServer()
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("ShutDown Server...")

	// 设置一个5s后返回Cancel的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// srv.Shutdown(ctx)的意思是:
	// 它需要完成两个任务：首先要保证所有已经建立的连接服务完毕
	// ctx的作用相当于超时定时器，当ctx返回Cancel时，srv必须结束，如果此时还有连接没完成，就会返回error
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("time out for 5 seconds.")
	}
	log.Println("Server exiting")
}
