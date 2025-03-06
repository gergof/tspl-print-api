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

func (a *App) RenderBody(w http.ResponseWriter, r *http.Request) {
	endpointName := chi.URLParam(r, "endpoint")

	var args map[string]string

	if err := json.NewDecoder(r.Body).Decode(&args); err != nil {
		jsonResponse(w, http.StatusBadRequest, "Body must be a valid JSON")
		return
	}

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
		"codeList": string(code),
	})
}
