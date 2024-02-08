package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
    	return
	}
	photo_id := ps.ByName("photoId")
	ctx.Logger.Info(photo_id)
	db := ctx.Database
	photoid, err := strconv.Atoi(photo_id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
    	return
	}
	err = db.DeletePhoto(photoid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
    	return
	}

	// If no error, send a success response
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	msg := []byte("Successfully deleted the image")
	n, err := w.Write(msg)
	if err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if n != len(msg) {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

}
