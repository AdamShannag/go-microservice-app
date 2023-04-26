package handler

import (
	"errors"
	"github.com/AdamShannag/toolkit/v2"
	"net/http"

	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "handler")))
	// ErrBadRequest error.
	ErrBadRequest = errors.New("Bad Request!")
	tools         = toolkit.Tools{}
)

func render(w http.ResponseWriter, body interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")

	switch v := body.(type) {
	case error:
		tools.ErrorJSON(w, v, status)
	case nil:
		// do nothing
	default:
		tools.WriteJSON(w, status, toolkit.JSONResponse{
			Message: "Success!",
			Error:   false,
			Data:    v,
		})
	}
}
