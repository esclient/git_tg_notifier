package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	BotToken string `mapstructure:"TG_BOT_TOKEN"`
	ChatID   int64  `mapstructure:"CHAT_ID"`
	ThreadID int64  `mapstructure:"THREAD_ID"`
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	}

	viper.AutomaticEnv()
	if err := viper.BindEnv("TG_BOT_TOKEN"); err != nil {
		log.Fatalf("viper.BindEnv TG_BOT_TOKEN error: %v", err)
	}

	if err := viper.BindEnv("CHAT_ID"); err != nil {
		log.Fatalf("viper.BindEnv CHAT_ID error: %v", err)
	}

	if err := viper.BindEnv("THREAD_ID"); err != nil {
		log.Fatalf("viper.BindEnv THREAD_ID error: %v", err)
	}

	tgBotToken := viper.GetString("TG_BOT_TOKEN")
	chatID := viper.GetInt64("CHAT_ID")
	threadID := viper.GetInt64("THREAD_ID")

	return &Config{
		BotToken: tgBotToken,
		ChatID:   chatID,
		ThreadID: threadID,
	}
}
