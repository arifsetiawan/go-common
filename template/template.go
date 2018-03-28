package template

import (
	"bytes"
	"text/template"
)

// Tprintf is
func Tprintf(tmpl string, data map[string]interface{}) string {
	t := template.Must(template.New("templated").Parse(tmpl))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		return ""
	}
	return buf.String()
}
