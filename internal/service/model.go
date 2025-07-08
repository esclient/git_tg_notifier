package service

const newCommitTemplate = `🚀 Новый коммит в репозитории проекта *{{.Repo}}*!  
👤 *Автор:* [{{.Pusher}}](tg://user?id={{.PusherID}})  
📝 *Изменения:*  
{{.CommitText}}  
🔗 [Открыть репозиторий]({{.RepoURL}})`
