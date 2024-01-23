package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to view profiles", http.StatusForbidden)
		return
	}
	name := ps.ByName("name")
	db := ctx.Database
	banned, err := db.AmIBanned(name, ctx.User.Username)
	if err != nil {
		http.Error(w, "You have to be logged in to view this", http.StatusForbidden)
		return
	}
	if banned {
		http.Error(w, "The user has banned you", http.StatusForbidden)
		return
	}
	ctx.Logger.Info("Right before get User details")
	details, err := db.GetUserDetails(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx.Logger.Info("Right after get User details")

	jsonResponse, err := json.Marshal(details)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
