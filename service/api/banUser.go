package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging
	username := ctx.User.Username
	toban := ps.ByName("name")
	db := ctx.Database
	err := db.AddBan(username, toban)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		ctx.Logger.Info("The error is")
		ctx.Logger.Info(err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	msg := []byte("Successfully banned the user")
	n, err := w.Write(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

	if n != len(msg) {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}
}
