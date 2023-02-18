package prompt

import (
	"bytes"
	"text/template"

	"github.com/wrk-grp/errnie"
)

type Executor struct {
	tmpl *template.Template
	buf  *bytes.Buffer
}

func NewExecutor(prompt string) *Executor {
	tmpl, err := template.New("prompt").Parse(prompt)
	errnie.Handles(err)

	return &Executor{tmpl, bytes.NewBuffer([]byte{})}
}

func (executor *Executor) Execute(builder *Builder) string {
	executor.tmpl.Execute(executor.buf, builder)
	return executor.buf.String()
}
