package service

import (
	"context"
	"sync"

	"github.com/esclient/git_tg_notifier/internal/github"
)

type TelegramClient interface {
	SendMessage(chatID, threadID int64, text string) error
}

type GithubClient interface {
	GetFirstPRInfo(ctx context.Context, repoFullName, sha string) (*github.PRInfo, error)
}

type Service struct {
	githubClient   GithubClient
	telegramClient TelegramClient
	chatID         int64
	threadID       int64
	members        map[string]int64

	mu             sync.RWMutex
	processedCache map[string]string
}

func NewService(telegramClient TelegramClient, githubClient GithubClient, chatID, threadID int64, members map[string]int64) *Service {
	return &Service{
		telegramClient: telegramClient,
		githubClient:   githubClient,
		chatID:         chatID,
		threadID:       threadID,
		members:        members,

		processedCache: make(map[string]string),
	}
}
