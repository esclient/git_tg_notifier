package main

import (
	"log"
	"net/http"

	"github.com/esclient/git_tg_notifier/internal/config"
	"github.com/esclient/git_tg_notifier/internal/github"
	"github.com/esclient/git_tg_notifier/internal/handler"
	"github.com/esclient/git_tg_notifier/internal/service"
	"github.com/esclient/git_tg_notifier/internal/telegram"
)

func main() {
	cfg := config.LoadConfig()

	tgClient := telegram.NewClient(cfg.TgBotToken)
	ghClient := github.NewClient(cfg.GithubToken)
	service := service.NewService(
		tgClient,
		ghClient,
		cfg.ChatID,
		cfg.ThreadID,
		cfg.Members,
	)
	handler := handler.NewHandler(service)

	mux := http.NewServeMux()
	handler.RegisterRoutes(mux)

	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8300", mux))
}
