package service

const newCommitTemplate = `🚀 Новый коммит в репозитории проекта *{{.Repo}}*!  
👤 *Автор:* [{{.Pusher}}](tg://user?id={{.PusherID}})  
🔀 {{if .branchURL}}*Ветка:* [{{.Branch}}]({{.branchURL}}){{else}}*Ветка:* {{.Branch}}{{end}}
📝 *Изменения:*  
{{.CommitText}}  
🔗 [Открыть репозиторий]({{.RepoURL}})`
