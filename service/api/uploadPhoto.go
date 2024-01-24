package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)

func uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Extract username from context
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to upload", http.StatusUnauthorized)
		return
	}
	username := ctx.User.Username
	ctx.Logger.Info("Called successfully")
	// Read image data from the request body
	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // For example, max 10 MB file size
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	// Retrieve the file from form data
	file, _, err := r.FormFile("image") // "image" should be the name of your file input field
	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file data
	ImageData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading image data", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	ctx.Logger.Info(ImageData)
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
