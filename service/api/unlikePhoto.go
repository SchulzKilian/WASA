package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Placeholder logic
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusForbidden)
		return
	}
	photoid := ps.ByName("photoId")
	err := ctx.Database.DeleteLike(ctx.User.Username, photoid)
	if err != nil {
		http.Error(w, "something went wrong with you trying to remove a like", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte("Successfully unliked the image"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
