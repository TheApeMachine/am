package prompt

import (
	"bytes"
	"text/template"

	"github.com/wrk-grp/errnie"
)

type Executor struct {
	tmpl *template.Template
}

func NewExecutor(prompt string) *Executor {
	errnie.Trace()

	tmpl, err := template.New("prompt").Parse(prompt)
	errnie.Handles(err)

	return &Executor{tmpl}
}

func (executor *Executor) Execute(builder *Builder) string {
	errnie.Trace()
	var buf bytes.Buffer
	errnie.Handles(executor.tmpl.Execute(&buf, builder))
	return buf.String()
}
