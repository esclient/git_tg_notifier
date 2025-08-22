package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	TgBotToken  string           `mapstructure:"TG_BOT_TOKEN"`
	GithubToken string           `mapstructure:"GITHUB_TOKEN"`
	ChatID      int64            `mapstructure:"CHAT_ID"`
	ThreadID    int64            `mapstructure:"THREAD_ID"`
	Members     map[string]int64 `mapstructure:"-"`
}

func LoadConfig() *Config {
	if _, err := os.Stat(".env"); err == nil {
		godotenv.Load()
	}

	viper.AutomaticEnv()
	if err := viper.BindEnv("TG_BOT_TOKEN"); err != nil {
		log.Fatalf("viper.BindEnv TG_BOT_TOKEN error: %v", err)
	}

	if err := viper.BindEnv("GITHUB_TOKEN"); err != nil {
		log.Fatalf("viper.BindEnv GITHUB_TOKEN error: %v", err)
	}

	if err := viper.BindEnv("CHAT_ID"); err != nil {
		log.Fatalf("viper.BindEnv CHAT_ID error: %v", err)
	}

	if err := viper.BindEnv("THREAD_ID"); err != nil {
		log.Fatalf("viper.BindEnv THREAD_ID error: %v", err)
	}

	tgBotToken := viper.GetString("TG_BOT_TOKEN")
	githubToken := viper.GetString("GITHUB_TOKEN")
	chatID := viper.GetInt64("CHAT_ID")
	threadID := viper.GetInt64("THREAD_ID")

	names := []string{"SAMU", "ANDR", "VAN", "LSH", "NKT", "JEN", "TIM"}
	members := make(map[string]int64, len(names))
	for _, key := range names {
		envNick := key + "_GITHUB_NICK"
		envID := key + "_TG_ID"

		if err := viper.BindEnv(envNick); err != nil {
			log.Fatalf("viper.BindEnv %s error: %v", envNick, err)
		}

		if err := viper.BindEnv(envID); err != nil {
			log.Fatalf("viper.BindEnv %s error: %v", envID, err)
		}

		members[viper.GetString(envNick)] = viper.GetInt64(envID)
	}

	return &Config{
		TgBotToken:  tgBotToken,
		GithubToken: githubToken,
		ChatID:      chatID,
		ThreadID:    threadID,
		Members:     members,
	}
}
