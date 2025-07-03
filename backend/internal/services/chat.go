package services

import (
	"context"
	"mi-amorito-backend/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/sashabaranov/go-openai"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChatService struct {
	chatCollection    *mongo.Collection
	sessionCollection *mongo.Collection
	openaiService     *OpenAIService
	characterService  *CharacterService
}

func NewChatService(db *mongo.Database, openaiService *OpenAIService, characterService *CharacterService) *ChatService {
	return &ChatService{
		chatCollection:    db.Collection("chats"),
		sessionCollection: db.Collection("chat_sessions"),
		openaiService:     openaiService,
		characterService:  characterService,
	}
}

// SendMessage 发送消息并获取AI回复
func (s *ChatService) SendMessage(userID primitive.ObjectID, characterID, content, sessionID string) (*models.SendMessageResponse, error) {
	// 获取角色信息
	character, err := s.characterService.GetCharacterByID(characterID)
	if err != nil {
		return nil, err
	}

	// 如果没有sessionID，创建新会话
	if sessionID == "" {
		sessionID = uuid.New().String()
		
		// 生成会话标题
		title, _ := s.openaiService.GenerateTitle(content)
		
		// 创建新会话
		err = s.createChatSession(userID, characterID, sessionID, title)
		if err != nil {
			return nil, err
		}
	}

	// 保存用户消息
	now := time.Now()
	userMessage := models.ChatMessage{
		ID:          primitive.NewObjectID(),
		SessionID:   sessionID,
		UserID:      userID,
		CharacterID: characterID,
		Role:        "user",
		Content:     content,
		Timestamp:   now,
		CreatedAt:   now,
	}

	err = s.saveChatMessage(userMessage)
	if err != nil {
		return nil, err
	}

	// 获取聊天历史
	chatHistory, err := s.getChatHistory(sessionID, 10)
	if err != nil {
		return nil, err
	}

	// 构建OpenAI消息历史
	var openaiMessages []openai.ChatCompletionMessage
	for _, msg := range chatHistory {
		role := openai.ChatMessageRoleUser
		if msg.Role == "assistant" {
			role = openai.ChatMessageRoleAssistant
		}
		
		openaiMessages = append(openaiMessages, openai.ChatCompletionMessage{
			Role:    role,
			Content: msg.Content,
		})
	}

	// 生成AI回复
	aiResponse, err := s.openaiService.GenerateResponse(character.SystemPrompt, content, openaiMessages)
	if err != nil {
		return nil, err
	}

	// 保存AI回复
	aiMessage := models.ChatMessage{
		ID:          primitive.NewObjectID(),
		SessionID:   sessionID,
		UserID:      userID,
		CharacterID: characterID,
		Role:        "assistant",
		Content:     aiResponse,
		Timestamp:   time.Now(),
		CreatedAt:   time.Now(),
	}

	err = s.saveChatMessage(aiMessage)
	if err != nil {
		return nil, err
	}

	// 更新会话最后消息
	err = s.updateSessionLastMessage(sessionID, aiResponse)
	if err != nil {
		// 不阻断流程，记录错误即可
	}

	return &models.SendMessageResponse{
		SessionID: sessionID,
		Message:   userMessage,
		Reply:     aiMessage,
	}, nil
}

// GetChatHistory 获取聊天历史
func (s *ChatService) GetChatHistory(sessionID string, limit int, skip int) (*models.GetChatHistoryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 构建查询条件
	filter := bson.M{"session_id": sessionID}
	opts := options.Find().
		SetSort(bson.M{"timestamp": 1}).
		SetLimit(int64(limit)).
		SetSkip(int64(skip))

	cursor, err := s.chatCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.ChatMessage
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	// 检查是否有更多消息
	totalCount, _ := s.chatCollection.CountDocuments(ctx, filter)
	hasMore := int64(skip+limit) < totalCount

	return &models.GetChatHistoryResponse{
		Messages: messages,
		HasMore:  hasMore,
	}, nil
}

// GetUserChatSessions 获取用户的聊天会话列表
func (s *ChatService) GetUserChatSessions(userID primitive.ObjectID) (*models.GetChatSessionsResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"user_id": userID}
	opts := options.Find().SetSort(bson.M{"updated_at": -1}).SetLimit(50)

	cursor, err := s.sessionCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var sessions []models.ChatSession
	if err = cursor.All(ctx, &sessions); err != nil {
		return nil, err
	}

	return &models.GetChatSessionsResponse{
		Sessions: sessions,
	}, nil
}

// 私有方法

func (s *ChatService) createChatSession(userID primitive.ObjectID, characterID, sessionID, title string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	now := time.Now()
	session := models.ChatSession{
		ID:          primitive.NewObjectID(),
		SessionID:   sessionID,
		UserID:      userID,
		CharacterID: characterID,
		Title:       title,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err := s.sessionCollection.InsertOne(ctx, session)
	return err
}

func (s *ChatService) saveChatMessage(message models.ChatMessage) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := s.chatCollection.InsertOne(ctx, message)
	return err
}

func (s *ChatService) getChatHistory(sessionID string, limit int) ([]models.ChatMessage, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"session_id": sessionID}
	opts := options.Find().
		SetSort(bson.M{"timestamp": -1}).
		SetLimit(int64(limit))

	cursor, err := s.chatCollection.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var messages []models.ChatMessage
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	// 反转顺序，使最旧的消息在前
	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func (s *ChatService) updateSessionLastMessage(sessionID, lastMessage string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"session_id": sessionID}
	update := bson.M{
		"$set": bson.M{
			"last_message": lastMessage,
			"updated_at":   time.Now(),
		},
	}

	_, err := s.sessionCollection.UpdateOne(ctx, filter, update)
	return err
}