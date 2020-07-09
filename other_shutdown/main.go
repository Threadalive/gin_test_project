package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		time.Sleep(11 * time.Second)
		c.String(200, "hello test")
	})

	server := &http.Server{
		Addr:    ":8085",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	//捕获两类信号
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	//阻塞channel
	<-quit
	log.Println("shutdown server...")

	//创建超时的上下文
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//执行关闭服务器
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown", err)
	}

	log.Println("server exiting...")
}
