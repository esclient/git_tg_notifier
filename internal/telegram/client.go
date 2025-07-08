package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TelegramClient struct {
	token      string
	apiURL     string
	httpClient *http.Client
}

func NewClient(token string) *TelegramClient {
	return &TelegramClient{
		token:      token,
		apiURL:     fmt.Sprintf("https://api.telegram.org/bot%s", token),
		httpClient: &http.Client{},
	}
}

func (c *TelegramClient) SendMessage(chatID, threadID int64, text string) error {
	message := TelegramMessage{
		ChatID:                chatID,
		ThreadID:              threadID,
		Text:                  text,
		ParseMode:             "Markdown",
		DisableWebPagePreview: 1,
	}

	data, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("json.Marshal error: %v", err)
	}

	resp, err := http.Post(
		c.apiURL+"/sendMessage",
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return fmt.Errorf("http.Post error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status: %s", resp.Status)
	}

	return nil
}
