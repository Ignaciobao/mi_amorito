package services

import (
	"context"
	"fmt"

	"github.com/sashabaranov/go-openai"
)

type OpenAIService struct {
	client *openai.Client
	model  string
}

func NewOpenAIService(apiKey, model string) *OpenAIService {
	client := openai.NewClient(apiKey)
	return &OpenAIService{
		client: client,
		model:  model,
	}
}

// GenerateResponse 生成角色回复
func (s *OpenAIService) GenerateResponse(systemPrompt, userMessage string, chatHistory []openai.ChatCompletionMessage) (string, error) {
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
	}

	// 添加聊天历史（最近10条消息）
	if len(chatHistory) > 10 {
		chatHistory = chatHistory[len(chatHistory)-10:]
	}
	messages = append(messages, chatHistory...)

	// 添加用户当前消息
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: userMessage,
	})

	req := openai.ChatCompletionRequest{
		Model:       s.model,
		Messages:    messages,
		MaxTokens:   500,
		Temperature: 0.8,
		TopP:        0.9,
	}

	resp, err := s.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "", fmt.Errorf("OpenAI API error: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from OpenAI")
	}

	return resp.Choices[0].Message.Content, nil
}

// GenerateTitle 为聊天会话生成标题
func (s *OpenAIService) GenerateTitle(firstMessage string) (string, error) {
	systemPrompt := "请根据用户的第一条消息，生成一个简短的聊天标题，不超过10个字符。只返回标题内容，不要其他说明。"
	
	messages := []openai.ChatCompletionMessage{
		{
			Role:    openai.ChatMessageRoleSystem,
			Content: systemPrompt,
		},
		{
			Role:    openai.ChatMessageRoleUser,
			Content: firstMessage,
		},
	}

	req := openai.ChatCompletionRequest{
		Model:       s.model,
		Messages:    messages,
		MaxTokens:   20,
		Temperature: 0.5,
	}

	resp, err := s.client.CreateChatCompletion(context.Background(), req)
	if err != nil {
		return "新的聊天", nil // 如果生成失败，返回默认标题
	}

	if len(resp.Choices) == 0 {
		return "新的聊天", nil
	}

	return resp.Choices[0].Message.Content, nil
}