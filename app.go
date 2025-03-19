package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type App struct {
	config *Config
}

func NewApp(config *Config) *App {
	return &App{
		config: config,
	}
}

func (a *App) Ping(w http.ResponseWriter, r *http.Request) {
	jsonResponse(w, http.StatusOK, "pong")
}

func (a *App) RenderWithParams(w http.ResponseWriter, r *http.Request) {
	args := make(map[string]string)

	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			args[key] = values[0]
		}
	}

	a.render(w, r, args)
}

func (a *App) RenderWithBody(w http.ResponseWriter, r *http.Request) {
	var args map[string]string

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Body must be a valid JSON")
		return
	}

	a.render(w, r, args)
}

func (a *App) render(w http.ResponseWriter, r *http.Request, args map[string]string) {
	endpointName := chi.URLParam(r, "endpoint")

	endpoint, exists := a.config.Endpoints[endpointName]
	if !exists {
		jsonResponse(w, http.StatusNotFound, fmt.Sprintf("Endpoint %s does not exist", endpointName))
		return
	}

	code, err := endpoint.RenderCodeList(endpoint.MapArgs(args))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Could not render code list: %v", err))
		return
	}

	jsonResponseObject(w, http.StatusOK, map[string]string{
		"message":  "CodeList rendered",
		"codeList": code,
	})
}

func (a *App) PrintWithParams(w http.ResponseWriter, r *http.Request) {
	args := make(map[string]string)

	for key, values := range r.URL.Query() {
		if len(values) > 0 {
			args[key] = values[0]
		}
	}

	a.print(w, r, args)
}

func (a *App) PrintWithBody(w http.ResponseWriter, r *http.Request) {
	var args map[string]string

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Body must be a valid JSON")
		return
	}

	a.print(w, r, args)
}

func (a *App) print(w http.ResponseWriter, r *http.Request, args map[string]string) {
	endpointName := chi.URLParam(r, "endpoint")

	endpoint, exists := a.config.Endpoints[endpointName]
	if !exists {
		jsonResponse(w, http.StatusNotFound, fmt.Sprintf("Endpoint %s does not exist", endpointName))
		return
	}

	code, err := endpoint.RenderCodeList(endpoint.MapArgs(args))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Could not render code list: %v", err))
		return
	}

	err = endpoint.Printer.SendCommand([]byte(code))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Failed to print: %v", err))
		return
	}

	jsonResponse(w, http.StatusOK, "Print requested")
}
