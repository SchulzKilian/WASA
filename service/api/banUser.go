package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusUnauthorized)
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging
	username := ctx.User.Username
	toban := ps.ByName("name")
	db := ctx.Database
	err := db.AddBan(username, toban)
	if err != nil {
		http.Error(w, "Error banning the user", http.StatusUnauthorized)
		ctx.Logger.Info(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	msg := []byte("Successfully banned the user")
	n, err := w.Write(msg)
	if err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if n != len(msg) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
