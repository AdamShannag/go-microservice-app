package main

import (
	"log"
	"net/http"

	"github.com/AdamShannag/toolkit/v2"
	"github.com/go-chi/chi"
)

func (app *Config) createUser(w http.ResponseWriter, req *http.Request) {
	user := app.Models.User

	err := app.tools.ReadJSON(w, req, &user)
	if err != nil {
		log.Println("an error has occured here: ", err)
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	_, err = user.Insert(user)
	if err != nil {
		log.Println("an error has occured here: ", err)
		app.tools.ErrorJSON(w, err)
		return
	}

	err = app.tools.WriteJSON(w, http.StatusOK, toolkit.JSONResponse{
		Message: "User added!",
		Error:   false,
	})

	if err != nil {
		log.Println("an error has occured here: ", err)
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) getUser(w http.ResponseWriter, req *http.Request) {
	user := app.Models.User

	id := chi.URLParam(req, "id")

	u, err := user.GetOne(id)
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
	err = app.tools.WriteJSON(w, http.StatusOK, toolkit.JSONResponse{
		Message: "Success!",
		Error:   false,
		Data:    u,
	})
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) getUsers(w http.ResponseWriter, req *http.Request) {
	user := app.Models.User

	users, err := user.GetAll()
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.tools.WriteJSON(w, http.StatusOK, toolkit.JSONResponse{
		Message: "Success!",
		Error:   false,
		Data:    users,
	})
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}

func (app *Config) updateUser(w http.ResponseWriter, req *http.Request) {
	user := app.Models.User

	err := app.tools.ReadJSON(w, req, &user)
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = user.Update()
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.tools.WriteJSON(w, http.StatusOK, toolkit.JSONResponse{
		Message: "User updated!",
		Error:   false,
	})

	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

}

func (app *Config) deleteUser(w http.ResponseWriter, req *http.Request) {
	user := app.Models.User

	id := chi.URLParam(req, "id")

	err := user.DeleteByID(id)
	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}

	err = app.tools.WriteJSON(w, http.StatusOK, toolkit.JSONResponse{
		Message: "User Deleted!",
		Error:   false,
	})

	if err != nil {
		app.tools.ErrorJSON(w, err, http.StatusBadRequest)
		return
	}
}
