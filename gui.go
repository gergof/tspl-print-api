package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"

	"github.com/go-chi/chi/v5"
)

type labelGUIField struct {
	Name  string
	Value string
}

type labelGUIData struct {
	Endpoint string
	Fields   []labelGUIField
	Message  string
	Error    string
}

var labelGUITemplate = template.Must(template.New("label-gui").Parse(`<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<title>Print {{ .Endpoint }}</title>
</head>
<body>
	<h1>Print {{ .Endpoint }}</h1>
	{{ if .Message }}<p>{{ .Message }}</p>{{ end }}
	{{ if .Error }}<p>{{ .Error }}</p>{{ end }}
	<form method="post">
		{{ range .Fields }}
		<p>
			<label for="{{ .Name }}">{{ .Name }}</label><br>
			<input id="{{ .Name }}" name="{{ .Name }}" value="{{ .Value }}">
		</p>
		{{ end }}
		<button type="submit">Print</button>
	</form>
</body>
</html>
`))

func (a *App) LabelGUI(w http.ResponseWriter, r *http.Request) {
	endpointName := chi.URLParam(r, "endpoint")

	endpoint, exists := a.config.Endpoints[endpointName]
	if !exists {
		http.Error(w, fmt.Sprintf("Endpoint %s does not exist", endpointName), http.StatusNotFound)
		return
	}

	data := labelGUIData{
		Endpoint: endpointName,
		Fields:   endpointGUIFields(endpoint, nil),
	}

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			data.Error = fmt.Sprintf("Failed to parse form: %v", err)
			renderLabelGUI(w, data)
			return
		}

		args := endpoint.GetArgsFromForm(r.PostForm)
		data.Fields = endpointGUIFields(endpoint, r.PostForm)

		code, err := endpoint.RenderCodeList(args)
		if err != nil {
			data.Error = fmt.Sprintf("Could not render code list: %v", err)
			renderLabelGUI(w, data)
			return
		}

		if err := endpoint.Printer.SendCommand([]byte(code)); err != nil {
			data.Error = fmt.Sprintf("Failed to print: %v", err)
			renderLabelGUI(w, data)
			return
		}

		data.Message = "Print requested"
	}

	renderLabelGUI(w, data)
}

func renderLabelGUI(w http.ResponseWriter, data labelGUIData) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := labelGUITemplate.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Failed to render GUI: %v", err), http.StatusInternalServerError)
	}
}

func endpointGUIFields(endpoint Endpoint, values url.Values) []labelGUIField {
	names := make([]string, 0, len(endpoint.Args))
	for name := range endpoint.Args {
		names = append(names, name)
	}
	sort.Strings(names)

	fields := make([]labelGUIField, 0, len(names))
	for _, name := range names {
		fields = append(fields, labelGUIField{
			Name:  name,
			Value: values.Get(name),
		})
	}

	return fields
}
