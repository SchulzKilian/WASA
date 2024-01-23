package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if ctx.User == nil {
		http.Error(w, "You have to be logged in to like", http.StatusForbidden)
		return
	}
	photoid := ps.ByName("photoId")
	err := ctx.Database.AddLike(ctx.User.Username, photoid)
	if err != nil {
		http.Error(w, "something went wrong with you trying to like a photo", http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Successfully liked the image"))
	if err != nil {

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

}
