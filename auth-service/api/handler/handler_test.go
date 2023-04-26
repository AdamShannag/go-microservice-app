package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRender_render(t *testing.T) {
	tests := []struct {
		name     string
		data     interface{}
		response string
	}{
		{
			name:     "message",
			data:     "lorem",
			response: `{"data":"lorem", "error":false, "message":"Success!"}`,
		},
		{
			name:     "error",
			data:     errors.New("system error"),
			response: `{"error":true, "message":"system error"}`,
		},
		{
			name:     "nil",
			data:     nil,
			response: ``,
		},
		{
			name: "struct",
			data: struct {
				ID int `json:"id"`
			}{ID: 1},
			response: `{"data":{"id":1}, "error":false, "message":"Success!"}`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var (
				rr = httptest.NewRecorder()
			)

			render(rr, test.data, 200)
			if test.response != "" {
				assert.JSONEq(t, test.response, rr.Body.String())
			} else {
				assert.Equal(t, test.response, rr.Body.String())
			}
		})
	}
}
