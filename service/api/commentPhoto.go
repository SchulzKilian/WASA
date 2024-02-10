package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	commenter := ctx.User.Username
	photoid := ps.ByName("photoId")
	var comment database.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}

	// Assign commenter and photoId to the comment struct
	comment.Commenter = commenter
	comment.PhotoID = photoid

	// Process the comment (e.g., store in database)
	err = ctx.Database.AddComment(comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	msg := []byte("Successfully commented the image")
	n, err := w.Write(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

	if n != len(msg) {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

}

func getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	photoid := ps.ByName("photoId")
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	// Check if ctx.Database is nil
	if ctx.Database == nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

	comments, err := ctx.Database.GetComments(photoid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}

	// If comments is nil, initialize it as an empty slice
	if comments == nil {
		comments = []database.Comment{}
	}

	// Marshal the comments slice to JSON
	jsonResponse, err := json.Marshal(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return
	}

	// Write the JSON response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_,err = w.Write(jsonResponse)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) 
		return
	}

}
