package telegram

type TelegramMessage struct {
	ChatID                int64  `json:"chat_id"`
	ThreadID              int64  `json:"message_thread_id"`
	Text                  string `json:"text"`
	ParseMode             string `json:"parse_mode"`
	DisableWebPagePreview int64  `json:"disable_web_page_preview"`
}
