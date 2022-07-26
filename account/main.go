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
	// 可以在这里插入其他的记录器，以进行结构化或分层的日志记录
	log.Println("Starting server...")

	// 初始化数据库
	ds, err := initDS()
	if err != nil {
		log.Fatalf("Unable to initialize data sources: %v\n", err)
	}

	// router := gin.Default()
	router, err := inject(ds)
	if err != nil {
		log.Fatalf("Failure to inject data sources: %v\n", err)
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	// 优雅的服务器关闭 - https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/server.go
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to initialize server: %v\n", err)
		}
	}()

	log.Printf("Listening on port %v\n", srv.Addr)

	// 等待终止信号的通道
	quit := make(chan os.Signal)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// 这将阻塞，直到一个信号被传递到 quit 通道
	<-quit

	// 上下文用于通知服务器它有5秒钟的时间来完成当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭数据库
	if err := ds.Close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down data sources: %v\n", err)
	}

	// 关闭服务器
	log.Println("Shutting down server...")
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v\n", err)
	}
}
