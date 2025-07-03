package services

import (
	"context"
	"mi-amorito-backend/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserService struct {
	collection *mongo.Collection
}

func NewUserService(db *mongo.Database) *UserService {
	return &UserService{
		collection: db.Collection("users"),
	}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(deviceID, nickname string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 检查用户是否已存在
	existingUser, err := s.GetUserByDeviceID(deviceID)
	if err == nil && existingUser != nil {
		return existingUser, nil // 用户已存在，直接返回
	}

	// 创建新用户
	now := time.Now()
	user := models.User{
		ID:        primitive.NewObjectID(),
		DeviceID:  deviceID,
		Nickname:  nickname,
		Avatar:    "/avatars/default.jpg", // 默认头像
		CreatedAt: now,
		UpdatedAt: now,
	}

	_, err = s.collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUserByDeviceID 根据设备ID获取用户
func (s *UserService) GetUserByDeviceID(deviceID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"device_id": deviceID}
	
	var user models.User
	err := s.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // 用户不存在
		}
		return nil, err
	}

	return &user, nil
}

// GetUserByID 根据用户ID获取用户
func (s *UserService) GetUserByID(userID primitive.ObjectID) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"_id": userID}
	
	var user models.User
	err := s.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(userID primitive.ObjectID, updates bson.M) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	updates["updated_at"] = time.Now()
	
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": updates}
	
	opts := options.Update().SetUpsert(false)
	_, err := s.collection.UpdateOne(ctx, filter, update, opts)
	
	return err
}