package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User 用户模型
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	DeviceID  string             `bson:"device_id" json:"device_id"`
	Nickname  string             `bson:"nickname" json:"nickname"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// Character 角色模型
type Character struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CharacterID   string             `bson:"character_id" json:"character_id"`
	Name          string             `bson:"name" json:"name"`
	Avatar        string             `bson:"avatar" json:"avatar"`
	Description   string             `bson:"description" json:"description"`
	Personality   string             `bson:"personality" json:"personality"`
	Background    string             `bson:"background" json:"background"`
	SystemPrompt  string             `bson:"system_prompt" json:"system_prompt"`
	Greeting      string             `bson:"greeting" json:"greeting"`
	Tags          []string           `bson:"tags" json:"tags"`
	IsActive      bool               `bson:"is_active" json:"is_active"`
	SortOrder     int                `bson:"sort_order" json:"sort_order"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

// ChatMessage 聊天消息模型
type ChatMessage struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SessionID   string             `bson:"session_id" json:"session_id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CharacterID string             `bson:"character_id" json:"character_id"`
	Role        string             `bson:"role" json:"role"` // "user" 或 "assistant"
	Content     string             `bson:"content" json:"content"`
	Timestamp   time.Time          `bson:"timestamp" json:"timestamp"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
}

// ChatSession 聊天会话模型
type ChatSession struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	SessionID   string             `bson:"session_id" json:"session_id"`
	UserID      primitive.ObjectID `bson:"user_id" json:"user_id"`
	CharacterID string             `bson:"character_id" json:"character_id"`
	Title       string             `bson:"title" json:"title"`
	LastMessage string             `bson:"last_message" json:"last_message"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// API 请求和响应模型

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	DeviceID string `json:"device_id" binding:"required"`
	Nickname string `json:"nickname"`
}

// SendMessageRequest 发送消息请求
type SendMessageRequest struct {
	SessionID   string `json:"session_id"`
	CharacterID string `json:"character_id" binding:"required"`
	Content     string `json:"content" binding:"required"`
}

// SendMessageResponse 发送消息响应
type SendMessageResponse struct {
	SessionID string      `json:"session_id"`
	Message   ChatMessage `json:"message"`
	Reply     ChatMessage `json:"reply"`
}

// GetCharactersResponse 获取角色列表响应
type GetCharactersResponse struct {
	Characters []Character `json:"characters"`
}

// GetChatHistoryResponse 获取聊天历史响应
type GetChatHistoryResponse struct {
	Messages []ChatMessage `json:"messages"`
	HasMore  bool          `json:"has_more"`
}

// GetChatSessionsResponse 获取聊天会话列表响应
type GetChatSessionsResponse struct {
	Sessions []ChatSession `json:"sessions"`
}