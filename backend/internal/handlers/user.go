package handlers

import (
	"mi-amorito-backend/internal/models"
	"mi-amorito-backend/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request format",
			"code":  "INVALID_REQUEST",
		})
		return
	}

	user, err := h.userService.CreateUser(req.DeviceID, req.Nickname)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to create user",
			"code":  "USER_CREATION_FAILED",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

// GetProfile 获取用户资料
func (h *UserHandler) GetProfile(c *gin.Context) {
	deviceID, exists := c.Get("device_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "device ID not found",
			"code":  "DEVICE_ID_NOT_FOUND",
		})
		return
	}

	user, err := h.userService.GetUserByDeviceID(deviceID.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user profile",
			"code":  "USER_PROFILE_FAILED",
		})
		return
	}

	if user == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
			"code":  "USER_NOT_FOUND",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}