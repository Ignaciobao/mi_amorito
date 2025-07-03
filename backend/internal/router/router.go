package router

import (
	"mi-amorito-backend/internal/config"
	"mi-amorito-backend/internal/handlers"
	"mi-amorito-backend/internal/middleware"
	"mi-amorito-backend/internal/services"
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Setup(cfg *config.Config, userService *services.UserService, characterService *services.CharacterService, chatService *services.ChatService) *gin.Engine {
	r := gin.New()

	// 添加中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// CORS 配置
	origins := strings.Split(cfg.AllowedOrigins, ",")
	r.Use(cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "X-Device-ID"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ok",
			"timestamp": time.Now().Unix(),
		})
	})

	// 静态文件服务 (头像等)
	r.Static("/avatars", "./assets/avatars")

	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)
	characterHandler := handlers.NewCharacterHandler(characterService)
	chatHandler := handlers.NewChatHandler(chatService, userService)

	// API 路由组
	api := r.Group("/api/v1")
	{
		// 用户相关
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/profile", middleware.DeviceAuth(), userHandler.GetProfile)
		}

		// 角色相关
		characters := api.Group("/characters")
		{
			characters.GET("/", characterHandler.GetCharacters)
			characters.GET("/:character_id", characterHandler.GetCharacter)
		}

		// 聊天相关
		chat := api.Group("/chat")
		chat.Use(middleware.DeviceAuth())
		{
			chat.POST("/send", chatHandler.SendMessage)
			chat.GET("/history/:session_id", chatHandler.GetChatHistory)
			chat.GET("/sessions", chatHandler.GetChatSessions)
		}
	}

	return r
}