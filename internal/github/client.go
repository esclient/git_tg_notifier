package github

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/go-github/v60/github"
)

type GithubClient struct {
	gh *github.Client
}

func NewClient(token string) *GithubClient {
	return &GithubClient{
		gh: github.
			NewClient(http.DefaultClient).
			WithAuthToken(token),
	}
}

func (c *GithubClient) GetFirstPRInfo(ctx context.Context, repoFullName, sha string) (*PRInfo, error) {
	owner, repo, _ := strings.Cut(repoFullName, "/")
	prs, _, err := c.gh.PullRequests.ListPullRequestsWithCommit(
		ctx,
		owner,
		repo,
		sha,
		&github.ListOptions{},
	)

	if err != nil {
		return nil, err
	}

	if len(prs) == 0 {
		return nil, nil
	}

	pr := prs[0]
	return &PRInfo{
		Number: pr.GetNumber(),
		Title:  pr.GetTitle(),
		URL:    pr.GetHTMLURL(),
	}, nil
}
