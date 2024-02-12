package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	// "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	db := ctx.Database
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	tooutput, err := db.GetFollowedUsersPhotos(ctx.User.Username)
	if err != nil {
		ctx.Logger.Info(err)
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Serialize the tooutput slice to JSON
	jsonResponse, err := json.Marshal(tooutput)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

	// Write the JSON response
	w.WriteHeader(http.StatusOK)

	_, err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

}
