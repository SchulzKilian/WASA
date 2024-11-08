package api

import (
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	if ctx.User == nil {
		w.WriteHeader(http.StatusUnauthorized) // Sets the status code only
		return
	}
	photoid := ps.ByName("photoId")
	err := ctx.Database.AddLike(ctx.User.Username, photoid)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest) // Sets the status code only
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	_, err = w.Write([]byte("Successfully liked the image"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError) // Sets the status code only
		return
	}

}
