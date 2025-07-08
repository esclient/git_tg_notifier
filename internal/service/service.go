package service

type TelegramClient interface {
	SendMessage(chatID, threadID int64, text string) error
}

type Service struct {
	telegramClient TelegramClient
	chatID         int64
	threadID       int64
}

func NewService(telegramClient TelegramClient, chatID, threadID int64) *Service {
	return &Service{
		telegramClient: telegramClient,
		chatID:         chatID,
		threadID:       threadID,
	}
}
