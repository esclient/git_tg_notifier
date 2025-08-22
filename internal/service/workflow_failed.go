package service

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"log"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (s *Service) WorkflowFailed(event model.WorkflowFailedEvent) error {
	var (
		showPR  bool
		prURL   string
		prTitle string
	)

	repoFullName := event.Repository.FullName
	branch := event.WorkflowJob.HeadBranch
	commitSHA := event.WorkflowJob.HeadSHA

	if branch != "main" {
		if info, err := s.githubClient.GetFirstPRInfo(context.Background(), repoFullName, commitSHA); err == nil && info != nil {
			prURL = info.URL
			prTitle = info.Title
			showPR = true
		}
	}

	tmpl, _ := template.New("workflowFailed").Parse(workflowFailedTemplate)

	var messageBuf bytes.Buffer
	_ = tmpl.Execute(&messageBuf, map[string]interface{}{
		"Project":     repoFullName,
		"ShowPR":      showPR,
		"PRTitle":     prTitle,
		"PRURL":       prURL,
		"Branch":      branch,
		"JobName":     event.WorkflowJob.Name,
		"Sender":      s.getMention(event.Sender.Login),
		"RepoURL":     fmt.Sprintf("https://github.com/%s", event.Repository.FullName),
		"WorkflowURL": event.WorkflowJob.HTMLURL,
	})

	if err := s.telegramClient.SendMessage(s.chatID, s.threadID, messageBuf.String()); err != nil {
		log.Printf("telegramClient.SendMessage error: %v", err)
		return err
	}

	return nil
}
