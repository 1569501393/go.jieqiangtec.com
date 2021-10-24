package main

import (
	"context"
	"gin_test/internal/container"
	"gin_test/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// func main() {
// 	// 创建一个默认的路由引擎
// 	r := gin.Default()
// 	// GET：请求方式；/hello：请求的路径
// 	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
// 	r.GET("/hello", func(c *gin.Context) {
// 		// c.JSON：返回JSON格式的数据
// 		c.JSON(200, gin.H{
// 			"message": "Hello world!",
// 		})
// 	})
// 	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
// 	err := r.Run(":8888")
// 	if err != nil {
// 		return
// 	}
// }

// // 优雅地重启或停止
// func main() {
// 	router := gin.Default()
// 	// gin.SetMode("debug")
// 	router.GET("/", func(c *gin.Context) {
// 		// time.Sleep(5 * time.Second)
// 		c.String(http.StatusOK, "/Start")
// 	})
//
// 	router.GET("/hello", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello Gin")
// 	})
//
// 	router.GET("/test", v1.Test)
// 	// router.GET("/user", controller.User)
//
// 	srv := &http.Server{
// 		Addr: ":8081",
// 		// Addr:    ":8888",
// 		Handler: router,
// 	}
//
// 	go func() {
// 		// 服务连接
// 		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
// 			log.Fatalf("listen: %s \n", err)
// 		}
// 	}()
//
// 	// 等待中断星号以优雅地关闭服务器（设置 5 秒的超时时间）
// 	quit := make(chan os.Signal)
// 	signal.Notify(quit, os.Interrupt)
// 	<-quit
// 	log.Println("Shutdown Server ...")
//
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	defer cancel()
// 	if err := srv.Shutdown(ctx); err != nil {
// 		log.Fatal("Server shutdown", err)
// 	}
//
// 	log.Println("Server exiting")
// }

// 优雅地重启或停止
func main() {

	ctn := container.New()
	// router.GET("/user", controller.User)

	srv := &http.Server{
		Addr: ":8081",
		// Addr:    ":8888",
		Handler: router.New(ctn).Registry(),
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s \n", err)
		}
	}()

	// 等待中断星号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server shutdown", err)
	}

	log.Println("Server exiting")
}
