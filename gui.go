package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"sort"
	"unicode/utf8"

	"github.com/go-chi/chi/v5"
)

const maxEndpointNameLength = 20

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
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<title>Print {{ .Endpoint }}</title>
	<style>
		* {
			box-sizing: border-box;
		}

		body {
			max-width: 560px;
			margin: 0 auto;
			padding: 24px 16px;
			color: #202124;
			background: #f8f9fa;
			font-family: system-ui, -apple-system, BlinkMacSystemFont, "Segoe UI", sans-serif;
			line-height: 1.35;
		}

		h1 {
			margin: 0 0 20px;
			overflow-wrap: anywhere;
			font-size: 24px;
			font-weight: 600;
		}

		p {
			margin: 0 0 14px;
		}

		label {
			display: block;
			margin-bottom: 4px;
			font-size: 13px;
			font-weight: 600;
		}

		input,
		button {
			font: inherit;
		}

		input {
			width: 100%;
			min-height: 38px;
			padding: 6px 10px;
			border: 1px solid #b8c0cc;
			border-radius: 4px;
			background: #fff;
		}

		button {
			min-height: 38px;
			padding: 6px 14px;
			border: 1px solid #2f6feb;
			border-radius: 4px;
			color: #fff;
			background: #1a73e8;
			cursor: pointer;
		}

		@media (max-width: 420px) {
			body {
				padding: 16px 12px;
			}

			h1 {
				font-size: 20px;
			}

			button {
				width: 100%;
			}
		}
	</style>
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
		Endpoint: truncateEndpointName(endpointName),
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

func truncateEndpointName(name string) string {
	if utf8.RuneCountInString(name) <= maxEndpointNameLength {
		return name
	}

	runes := []rune(name)
	return string(runes[:maxEndpointNameLength])
}
