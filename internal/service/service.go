package service

import "sync"

type TelegramClient interface {
	SendMessage(chatID, threadID int64, text string) error
}

type Service struct {
	telegramClient TelegramClient
	chatID         int64
	threadID       int64
	members        map[string]int64

	mu             sync.RWMutex
	processedCache map[string]string
}

func NewService(telegramClient TelegramClient, chatID, threadID int64, members map[string]int64) *Service {
	return &Service{
		telegramClient: telegramClient,
		chatID:         chatID,
		threadID:       threadID,
		members:        members,

		processedCache: make(map[string]string),
	}
}
