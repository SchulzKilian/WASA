package api

import (
	"fmt"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	ctx.Logger.Info("myApiHandler called") // Example logging

	name := ps.ByName("name")
	if ctx.User != nil {
		err := ctx.Database.SetName(name, ctx.User.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

	} else {
		http.Error(w, "please log in first", http.StatusUnauthorized)
		return
	}
	err := ctx.Database.SetName(name, ctx.User.UserID)
	if err != nil {
		http.Error(w, "error changing the name in the database", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully changed your name.")
}
