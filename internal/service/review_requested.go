package service

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"sort"
	"strings"

	"github.com/esclient/git_tg_notifier/internal/model"
)

func (s *Service) ReviewRequested(event model.ReviewRequestedEvent) error {
	logins := make([]string, len(event.PullRequest.RequestedReviewers))
	for i, rv := range event.PullRequest.RequestedReviewers {
		logins[i] = rv.Login
	}
	sort.Strings(logins)

	var links []string
	for _, login := range logins {
		links = append(links, s.getMention(login))
	}

	reviewers := strings.Join(links, ", ")

	key := fmt.Sprintf("%s#%d",
		event.Repository.FullName,
		event.PullRequest.Number,
	)

	s.mu.RLock()
	prev, seen := s.processedCache[key]
	s.mu.RUnlock()
	if seen && prev == reviewers {
		return nil
	}

	tmpl, _ := template.New("reviewRequested").Parse(reviewRequestedTemplate)

	var messageBuf bytes.Buffer
	tmpl.Execute(&messageBuf, map[string]interface{}{
		"Author":    s.getMention(event.PullRequest.User.Login),
		"Title":     event.PullRequest.Title,
		"Reviewers": reviewers,
		"RepoName":  event.Repository.FullName,
		"RepoURL":   fmt.Sprintf("https://github.com/%s", event.Repository.FullName),
		"PRURL":     fmt.Sprintf("https://github.com/%s/pull/%d", event.Repository.FullName, event.PullRequest.Number),
	})

	if err := s.telegramClient.SendMessage(s.chatID, s.threadID, messageBuf.String()); err != nil {
		log.Printf("telegramClient.SendMessage error: %v", err)

		return err
	}

	s.mu.Lock()
	s.processedCache[key] = reviewers
	s.mu.Unlock()

	return nil
}
