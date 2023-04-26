package api

import (
	"github.com/AdamShannag/go-microservice-app/auth-service/api/handler"
	"github.com/AdamShannag/go-microservice-app/auth-service/services/users"
	"github.com/go-chi/chi"
	chimid "github.com/go-chi/chi/middleware"
	"github.com/go-rel/rel"
	"github.com/goware/cors"
)

// NewMux api.
func NewMux(repository rel.Repository) *chi.Mux {

	var (
		mux           = chi.NewMux()
		userService   = users.New(repository)
		healthHandler = handler.NewHealth()
		userHandler   = handler.NewUsers(repository, userService)
	)

	healthHandler.Add("database", repository)

	mux.Use(chimid.RequestID)
	mux.Use(chimid.RealIP)
	mux.Use(chimid.Recoverer)
	mux.Use(cors.AllowAll().Handler)

	mux.Mount("/actuator/health", healthHandler)
	mux.Mount("/users", userHandler)

	return mux
}
