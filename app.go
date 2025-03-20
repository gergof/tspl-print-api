package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/tidwall/gjson"
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

func (a *App) Render(w http.ResponseWriter, r *http.Request) {
	body, err := a.readBody(r)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Failed to read POST body")
	}

	if !gjson.Valid(body) {
		jsonResponse(w, http.StatusBadRequest, "Body must be a valid JSON")
		return
	}

	endpointName := chi.URLParam(r, "endpoint")

	endpoint, exists := a.config.Endpoints[endpointName]
	if !exists {
		jsonResponse(w, http.StatusNotFound, fmt.Sprintf("Endpoint %s does not exist", endpointName))
		return
	}

	code, err := endpoint.RenderCodeList(endpoint.GetArgsFromJson(body))
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, fmt.Sprintf("Could not render code list: %v", err))
		return
	}

	jsonResponseObject(w, http.StatusOK, map[string]string{
		"message":  "CodeList rendered",
		"codeList": code,
	})
}

func (a *App) Print(w http.ResponseWriter, r *http.Request) {
	body, err := a.readBody(r)
	if err != nil {
		jsonResponse(w, http.StatusInternalServerError, "Failed to read POST body")
	}

	if !gjson.Valid(body) {
		jsonResponse(w, http.StatusBadRequest, "Body must be a valid JSON")
		return
	}

	endpointName := chi.URLParam(r, "endpoint")

	endpoint, exists := a.config.Endpoints[endpointName]
	if !exists {
		jsonResponse(w, http.StatusNotFound, fmt.Sprintf("Endpoint %s does not exist", endpointName))
		return
	}

	code, err := endpoint.RenderCodeList(endpoint.GetArgsFromJson(body))
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

func (a *App) readBody(r *http.Request) (string, error) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return "", err
	}

	defer r.Body.Close()

	return string(body), nil
}
