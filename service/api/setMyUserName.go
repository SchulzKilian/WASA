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
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

	} else {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	err := ctx.Database.SetName(name, ctx.User.UserID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Successfully changed your name.")
}
