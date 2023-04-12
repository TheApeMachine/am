package prompt

import (
	"bytes"
	"text/template"

	"github.com/wrk-grp/errnie"
)

type Builder struct {
	System  string
	Current string
	Q       []map[string]string
}

func NewBuilder(system, current string) *Builder {
	errnie.Trace()
	return &Builder{system, current, make([]map[string]string, 0)}
}

func (builder *Builder) String() string {
	tmpl, err := template.New("prompt").Parse(builder.System + "\n" + builder.Current)
	errnie.Handles(err)
	var buf bytes.Buffer

	errnie.Handles(tmpl.Execute(&buf, builder))
	return buf.String()
}
