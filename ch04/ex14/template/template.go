package template

import (
	"html/template"
	"io"
	"log"

	"../github"
)

const templ = `<h1>{{.TotalCount}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>Title</th>
<th>State</th>
<th>User</th>
<th>MileStone</th>
</tr>
{{range .Items}}
<tr>
<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
<td>{{.Title}}</td>
<td>{{.State}}</td>
<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
{{if .MileStone}}
<td><a href='{{.MileStone.HTMLURL}}'>{{.MileStone.Title}}</a></td>
{{else}}
<td>none</td>
{{end}}
</tr>
{{end}}
</table>
`

var report *template.Template

func init() {
	report = template.Must(template.New("issuelist").Parse(templ))
}

// OutputHTML html出力
func OutputHTML(wr io.Writer, result *github.IssuesSearchResult) {
	if err := report.Execute(wr, result); err != nil {
		log.Fatal(err)
	}
}
