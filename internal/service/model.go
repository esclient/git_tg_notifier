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

const workflowFailedTemplate = `🚨 Ошибка в GitHub Actions! 🚨
👤 *Запустил*: {{.Sender}}
🔧 *Репозиторий*: [{{.Project}}]({{.RepoURL}})
{{- if .ShowPR }}
🔀 *Pull Request*: [{{.PRTitle}}]({{.PRURL}})
{{- else }}
🏷️ *Ветка*: {{.Branch}}
{{- end }}
⚠️ *Проблемная джоба*: ` + "`" + `{{.JobName}}` + "`" + `
🔗 [Открыть Workflow]({{.WorkflowURL}})
`
