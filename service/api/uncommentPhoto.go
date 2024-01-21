package api

import (

    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic

    id := ps.ByName("commentId")
    err := ctx.Database.DeleteComment(id, ctx.User.Username)
    if err != nil {
        http.Error(w, err.Error(), http.StatusForbidden)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Successfully uncommented the image"))
}