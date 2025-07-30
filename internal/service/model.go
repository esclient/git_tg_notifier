package service

const newCommitTemplate = `🚀 Новый коммит в репозитории *{{.Repo}}*!  
👤 *Автор:* {{.Author}} 
🔀 {{if .branchURL}}*Ветка:* [{{.Branch}}]({{.branchURL}}){{else}}*Ветка:* {{.Branch}}{{end}}
📝 *Изменения:*  
{{.CommitText}}  
🔗 [Открыть репозиторий]({{.RepoURL}})`

const reviewRequestedTemplate = `🔔 *Новый запрос ревью* 🚀  
👤 *Автор PR:* {{.Author}}
📦 *Репозиторий:* [{{.RepoName}}]({{.RepoURL}})
🏷️ *Название:* {{.Title}}
👥 *Ревьюверы:* {{.Reviewers}}
🔗 [Открыть PR]({{.PRURL}})`
