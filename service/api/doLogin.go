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
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	username, ok := requestData["name"]
	if !ok {
		http.Error(w, "Name field is required", http.StatusBadRequest)
		return
	}
	ctx.Logger.Info(username)
	db := ctx.Database
	userexists, err, token := db.DoesUserExist(username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking user existence")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info("Logged user in") // Example logging
	if userexists {
		fmt.Fprintf(w, token)
		return

	}
	if !userexists {

		var user database.User

		user.Username = username

		// Call AddUser method of the database object
		ctx.Logger.Info("Made it to the adduser call")

		err, token = db.AddUser(&user)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		ctx.Logger.Info("User created successfully")
		fmt.Fprintf(w, token)
		return
	}

}
