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
		http.Error(w, "You have to be logged in to view profiles", http.StatusUnauthorized)
		return
	}
	name := ps.ByName("name")
	db := ctx.Database
	banned, err := db.AmIBanned(name, ctx.User.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}
	if banned {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	ctx.Logger.Info("Right before get User details")
	details, err := db.GetUserDetails(name, ctx.User.Username)
	if err != nil {
		http.Error(w, err.Error(), http.StatusForbidden)
		return
	}
	ctx.Logger.Info("Right after get User details")

	jsonResponse, err := json.Marshal(details)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}
	ctx.Logger.Info("myApiHandler called") // Example logging

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

}
