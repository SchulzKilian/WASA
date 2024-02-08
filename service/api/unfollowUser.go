package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
    	return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging
	username := ctx.User.Username
	tofollow := ps.ByName("name")
	db := ctx.Database
	err := db.DeleteFollow(username, tofollow)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
    	return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte("Successfully unfollowed the user"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
    	return
	}

}
