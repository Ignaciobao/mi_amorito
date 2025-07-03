package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"mi-amorito-backend/internal/config"
	"mi-amorito-backend/internal/database"
	"mi-amorito-backend/internal/router"
	"mi-amorito-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg := config.Load()

	// 设置 Gin 模式
	gin.SetMode(cfg.GinMode)

	// 初始化数据库连接
	db, err := database.InitMongoDB(cfg.MongoURI, cfg.MongoDatabase)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer database.DisconnectMongoDB()

	// 初始化 Redis
	redisClient, err := database.InitRedis(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	defer redisClient.Close()

	// 初始化服务
	openaiService := services.NewOpenAIService(cfg.OpenAIAPIKey, cfg.OpenAIModel)
	characterService := services.NewCharacterService(db)
	chatService := services.NewChatService(db, openaiService, characterService)
	userService := services.NewUserService(db)

	// 初始化路由
	r := router.Setup(cfg, userService, characterService, chatService)

	// 创建服务器
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	// 启动服务器
	go func() {
		log.Printf("Server starting on port %s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// 优雅关闭
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}