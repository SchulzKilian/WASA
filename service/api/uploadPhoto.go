package api

import (
	"io/ioutil"
	"net/http"
	"time"

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
)

func uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract username from context
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusForbidden)
		return
	}
	username := ctx.User.Username
	ctx.Logger.Info("Called successfully")
	// Read image data from the request body
	ImageData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading image data", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()
	ctx.Logger.Info("Managed reading request")
	// Set current time as Timestamp
	Timestamp := time.Now()

	// Create a Photo struct
	photo := database.Photo{
		Username:  username,
		ImageData: ImageData,
		Timestamp: Timestamp,
	}

	// Call AddPhoto method to insert the photo into the database
	err = ctx.Database.AddPhoto(photo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with success message
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Photo uploaded successfully"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
