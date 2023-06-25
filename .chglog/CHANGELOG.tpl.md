# Changelog

{{range .Versions}}
## {{.Tag.Name}} ({{datetime "2006-01-02"}})

{{range .CommitGroups}}
### {{.Title}}

{{range .Commits}}
- {{if .Scope}}**{{.Scope}}:** {{end}}{{.Subject}}
{{end}}

{{end}}
{{end}}
