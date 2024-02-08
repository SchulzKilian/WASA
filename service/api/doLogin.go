package api

import (
	"encoding/json"
	"fmt"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var requestData map[string]string
	ctx.Logger.Info("test")
	// Decode the JSON body into the map
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
    	return
	}
	defer r.Body.Close()

	username, ok := requestData["name"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
    	return
	}
	ctx.Logger.Info(username)
	db := ctx.Database
	userexists, err, token := db.DoesUserExist(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
    	return
	}
	ctx.Logger.Info("Logged user in") // Example logging
	if userexists {
		fmt.Fprint(w, token)
		return

	}
	if !userexists {

		var user database.User

		user.Username = username

		// Call AddUser method of the database object
		ctx.Logger.Info("Made it to the adduser call")

		err, token = db.AddUser(&user)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		ctx.Logger.Info("User created successfully")
		fmt.Fprint(w, token)
		return
	}

}
