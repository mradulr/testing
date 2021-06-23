# Changelog

## {{.ReleaseDate}}

{{range .Modules1}}
* ### [{{.Name}}]({{.RemoteUrl}})
  * {{.Diff.Content}}
{{end}}

## {{.ReleaseDate}}

{{range .Modules2}}
* ### [{{.Name}}]({{.RemoteUrl}})
  * {{.Diff.Content}}
{{end}}