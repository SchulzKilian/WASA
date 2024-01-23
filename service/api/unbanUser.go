package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusForbidden)
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging
	username := ctx.User.Username
	tounban := ps.ByName("name")
	db := ctx.Database
	err := db.DeleteBan(username, tounban)
	if err != nil {
		http.Error(w, "Error unbanning the user", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Successfully unbanned the user"))
	if err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
