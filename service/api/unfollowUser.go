package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusForbidden)
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging
	username := ctx.User.Username
	tofollow := ps.ByName("name")
	db := ctx.Database
	err := db.DeleteFollow(username, tofollow)
	if err != nil {
		http.Error(w, "Error unfollowing the user", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte("Successfully unfollowed the user"))
	if err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
