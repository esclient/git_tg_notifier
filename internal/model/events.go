package model

type CommitEvent struct {
	Commits    []Commit   `json:"commits"`
	Repository Repository `json:"repository"`
	Pusher     Pusher     `json:"pusher"`
	Ref        string     `json:"ref"`
}

type Commit struct {
	Message string `json:"message"`
	URL     string `json:"url"`
	Author  Author `json:"author"`
}

type Author struct {
	Username string `json:"username"`
}

type Repository struct {
	Name    string `json:"name"`
	HTMLURL string `json:"html_url"`
}

type Pusher struct {
	Name string `json:"name"`
}

type ReviewRequestedEvent struct {
	Action     string `json:"action"`
	Repository struct {
		FullName string `json:"full_name"`
	} `json:"repository"`
	PullRequest struct {
		Number int    `json:"number"`
		Title  string `json:"title"`
		User   struct {
			Login   string `json:"login"`
			HTMLURL string `json:"html_url"`
		} `json:"user"`
		RequestedReviewers []struct {
			Login string `json:"login"`
		} `json:"requested_reviewers"`
	} `json:"pull_request"`
}
