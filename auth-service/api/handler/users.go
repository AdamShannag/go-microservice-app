package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AdamShannag/go-microservice-app/auth-service/services/users"
	"github.com/go-chi/chi"
	"github.com/go-rel/rel"
	"go.uber.org/zap"
	"net/http"
)

type ctx int

const (
	bodyKey ctx = 0
	loadKey ctx = 1
)

type Users struct {
	*chi.Mux
	repository rel.Repository
	users      users.Service
}

// Index handle GET /.
func (u Users) Index(w http.ResponseWriter, r *http.Request) {
	var (
		ctx    = r.Context()
		query  = r.URL.Query()
		result []users.User
		filter = users.Filter{
			Name: query.Get("name"),
		}
	)

	err := u.users.GetAll(ctx, &result, filter)
	if err != nil {
		logger.Warn("error while get users", zap.Error(err))
		render(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	render(w, result, http.StatusOK)
}

// Create handle POST /
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user users.User
	)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, http.StatusBadRequest)
		return
	}

	if err := u.users.Create(ctx, &user); err != nil {
		render(w, err, http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Location", fmt.Sprint(r.RequestURI, "/", user.ID))
	render(w, user, http.StatusCreated)
}

// Show handle GET /{ID}
func (u Users) Show(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user = ctx.Value(loadKey).(users.User)
	)

	render(w, user, http.StatusOK)
}

// Update handle PATCH /{ID}
func (u Users) Update(w http.ResponseWriter, r *http.Request) {
	var (
		ctx     = r.Context()
		user    = ctx.Value(loadKey).(users.User)
		changes = rel.NewChangeset(&user)
	)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		logger.Warn("decode error", zap.Error(err))
		render(w, ErrBadRequest, http.StatusBadRequest)
		return
	}

	if err := u.users.Update(ctx, &user, changes); err != nil {
		render(w, err, http.StatusUnprocessableEntity)
		return
	}

	render(w, user, http.StatusOK)
}

// Destroy handle DELETE /{ID}
func (u Users) Destroy(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user = ctx.Value(loadKey).(users.User)
	)

	u.users.Delete(ctx, &user)
	//if err != nil {
	//	logger.Warn(fmt.Sprintf("error while deleting a user [%s]", user.ID), zap.Error(err))
	//	render(w, ErrBadRequest, http.StatusBadRequest)
	//	return
	//}
	render(w, nil, http.StatusNoContent)
}

// Load is middleware that loads todos to context.
func (u Users) Load(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			ctx  = r.Context()
			id   = chi.URLParam(r, "ID")
			user users.User
		)

		if err := u.users.Get(ctx, &user, id); err != nil {
			if errors.Is(err, rel.ErrNotFound) {
				render(w, err, http.StatusNotFound)
				return
			}
			panic(err)
		}

		ctx = context.WithValue(ctx, loadKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// NewUsers handler.
func NewUsers(repository rel.Repository, users users.Service) Users {
	h := Users{
		Mux:        chi.NewMux(),
		repository: repository,
		users:      users,
	}

	h.Get("/", h.Index)
	h.Post("/", h.Create)
	h.With(h.Load).Get("/{ID}", h.Show)
	h.With(h.Load).Put("/{ID}", h.Update)
	h.With(h.Load).Delete("/{ID}", h.Destroy)

	return h
}
