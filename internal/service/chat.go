package service

import "github.com/0xpelamar/chatbot/internal/repository"

type Chat struct {
	repository.Chat
}

func NewChatService(chat repository.Chat) *Chat {
	return &Chat{
		Chat: chat,
	}
}
