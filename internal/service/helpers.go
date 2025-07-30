package service

import "fmt"

func (s *Service) getTelegramID(githubNick string) (tgID int64, ok bool) {
	tgID, ok = s.members[githubNick]
	return
}

func (s *Service) getMention(githubNick string) string {
	if tgID, ok := s.getTelegramID(githubNick); ok {
		return fmt.Sprintf("[%s](tg://user?id=%d)", githubNick, tgID)
	}

	return fmt.Sprintf("[%s](https://github.com/%s)", githubNick, githubNick)
}
