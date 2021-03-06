package main

import (
	"html/template"
)

var s = `
<h1>{{.PageTitle}}<h1>
<ul>
    {{range .Todos}}
        {{if .Done}}
            <li class="done">{{.Title}}</li>
        {{else}}
            <li>{{.Title}}</li>
        {{end}}
    {{end}}
</ul>
`

var t *template.Template

func f() {

	// do task with template
}

func main() {
	// costs time at load and maybe unused
	t = template.Must(template.New("").Parse(s))
	_ = t

	for i := 0; i < 10000; i++ {
		f()
	}
}
