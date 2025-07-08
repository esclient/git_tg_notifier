package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"strings"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (s *Service) Commit(event model.CommitEvent) error {
	lastCommit := event.Commits[len(event.Commits)-1]
	title := lastCommit.Message
	if idx := strings.IndexByte(title, '\n'); idx != -1 {
		title = title[:idx]
	}

	tmpl, _ := template.New("commit").Parse(newCommitTemplate)

	var messageBuf bytes.Buffer
	tmpl.Execute(&messageBuf, map[string]interface{}{
		"Repo":       event.Repository.Name,
		"Pusher":     event.Pusher.Name,
		"PusherID":   event.Pusher.ID,
		"CommitText": fmt.Sprintf("- [%s](%s)", title, lastCommit.URL),
		"RepoURL":    event.Repository.HTMLURL,
	})

	if err := s.telegramClient.SendMessage(s.chatID, s.threadID, messageBuf.String()); err != nil {
		log.Printf("telegramClient.SendMessage error: %v", err)

		return err
	}

	return nil
}
