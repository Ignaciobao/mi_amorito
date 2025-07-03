package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitMongoDB(uri, database string) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 创建客户端选项
	clientOptions := options.Client().ApplyURI(uri)

	// 连接到 MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// 检查连接
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	mongoClient = client
	db := client.Database(database)

	// 创建索引
	if err := createIndexes(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

func DisconnectMongoDB() error {
	if mongoClient != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return mongoClient.Disconnect(ctx)
	}
	return nil
}

func createIndexes(ctx context.Context, db *mongo.Database) error {
	// 用户集合索引
	userCollection := db.Collection("users")
	userIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"device_id": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"created_at": 1},
		},
	}
	_, err := userCollection.Indexes().CreateMany(ctx, userIndexes)
	if err != nil {
		return err
	}

	// 聊天记录集合索引
	chatCollection := db.Collection("chats")
	chatIndexes := []mongo.IndexModel{
		{
			Keys: map[string]interface{}{"user_id": 1, "character_id": 1},
		},
		{
			Keys: map[string]interface{}{"user_id": 1, "created_at": -1},
		},
		{
			Keys: map[string]interface{}{"session_id": 1},
		},
	}
	_, err = chatCollection.Indexes().CreateMany(ctx, chatIndexes)
	if err != nil {
		return err
	}

	// 角色集合索引
	characterCollection := db.Collection("characters")
	characterIndexes := []mongo.IndexModel{
		{
			Keys:    map[string]interface{}{"character_id": 1},
			Options: options.Index().SetUnique(true),
		},
		{
			Keys: map[string]interface{}{"is_active": 1},
		},
	}
	_, err = characterCollection.Indexes().CreateMany(ctx, characterIndexes)
	return err
}