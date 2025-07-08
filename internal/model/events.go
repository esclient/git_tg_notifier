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
	ID   int64  `json:"id"`
}
