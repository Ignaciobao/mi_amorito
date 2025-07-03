package handlers

import (
	"mi-amorito-backend/internal/models"
	"mi-amorito-backend/internal/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct {
	chatService *services.ChatService
	userService *services.UserService
}

func NewChatHandler(chatService *services.ChatService, userService *services.UserService) *ChatHandler {
	return &ChatHandler{
		chatService: chatService,
		userService: userService,
	}
}

// SendMessage 发送消息
func (h *ChatHandler) SendMessage(c *gin.Context) {
	deviceID, exists := c.Get("device_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "device ID not found",
			"code":  "DEVICE_ID_NOT_FOUND",
		})
		return
	}

	// 获取用户信息
	user, err := h.userService.GetUserByDeviceID(deviceID.(string))
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
			"code":  "USER_NOT_FOUND",
		})
		return
	}

	var req models.SendMessageRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
			"code":  "INVALID_REQUEST",
		})
		return
	}

	// 发送消息并获取回复
	response, err := h.chatService.SendMessage(user.ID, req.CharacterID, req.Content, req.SessionID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to send message",
			"code":  "MESSAGE_SEND_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetChatHistory 获取聊天历史
func (h *ChatHandler) GetChatHistory(c *gin.Context) {
	sessionID := c.Param("session_id")
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "session ID is required",
			"code":  "SESSION_ID_REQUIRED",
		})
		return
	}

	// 获取分页参数
	limitStr := c.DefaultQuery("limit", "20")
	skipStr := c.DefaultQuery("skip", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 20
	}

	skip, err := strconv.Atoi(skipStr)
	if err != nil {
		skip = 0
	}

	// 限制每次查询的最大数量
	if limit > 100 {
		limit = 100
	}

	response, err := h.chatService.GetChatHistory(sessionID, limit, skip)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get chat history",
			"code":  "CHAT_HISTORY_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// GetChatSessions 获取用户的聊天会话列表
func (h *ChatHandler) GetChatSessions(c *gin.Context) {
	deviceID, exists := c.Get("device_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "device ID not found",
			"code":  "DEVICE_ID_NOT_FOUND",
		})
		return
	}

	// 获取用户信息
	user, err := h.userService.GetUserByDeviceID(deviceID.(string))
	if err != nil || user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user not found",
			"code":  "USER_NOT_FOUND",
		})
		return
	}

	response, err := h.chatService.GetUserChatSessions(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get chat sessions",
			"code":  "CHAT_SESSIONS_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, response)
}