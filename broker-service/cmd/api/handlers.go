package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	toolkit "github.com/AdamShannag/toolkit/v2"
)

type RequestPayload struct {
	Action string      `json:"action"`
	User   UserPayload `json:"user,omitempty"`
}

type UserPayload struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name"`
	Address string `json:"address,omitempty"`
}

func (app *Config) Broker(w http.ResponseWriter, r *http.Request) {
	payload := toolkit.JSONResponse{
		Error:   false,
		Message: "Broker working",
	}

	_ = app.tools.WriteJSON(w, http.StatusOK, payload)
}

func (app *Config) HandleSubmission(w http.ResponseWriter, r *http.Request) {
	var requestPayload RequestPayload
	err := app.tools.ReadJSON(w, r, &requestPayload)
	if err != nil {
		app.tools.ErrorJSON(w, err)
		return
	}

	switch requestPayload.Action {
	case "user":
		app.userRequest(w, requestPayload.User, r.URL.Query(), r.Method)
	default:
		app.tools.ErrorJSON(w, errors.New("unknown action"))
	}
}

func (app *Config) userRequest(w http.ResponseWriter, user UserPayload, params url.Values, method string) {
	jsonData, err := json.Marshal(user)

	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	userServiceUrl := USER_SERVICE_URL
	if params.Has("id") {
		userServiceUrl = fmt.Sprintf("%s/%s", userServiceUrl, params.Get("id"))
	}
	request, err := http.NewRequest(method, userServiceUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		app.tools.ErrorJSON(w, err)
		return
	}

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		app.tools.ErrorJSON(w, err)
		return
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		app.tools.ErrorJSON(w, errors.New("something went wrong"))
		return
	}

	var jsonFromService toolkit.JSONResponse

	err = json.NewDecoder(response.Body).Decode(&jsonFromService)
	if err != nil {
		app.tools.ErrorJSON(w, err)
		return
	}

	if jsonFromService.Error {
		app.tools.ErrorJSON(w, err)
		return
	}

	app.tools.WriteJSON(w, http.StatusOK, jsonFromService)
}
