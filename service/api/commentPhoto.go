package api

import (
    "net/http"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    commenter := ctx.User.Username
    photoid := ps.ByName("photoId")
    var comment database.Comment
    err := json.NewDecoder(r.Body).Decode(&comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Assign commenter and photoId to the comment struct
    comment.Commenter = commenter
    comment.PhotoID = photoid

    // Process the comment (e.g., store in database)
    err = ctx.Database.AddComment(comment)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    msg := []byte("Successfully commented the image")
    n, err := w.Write(msg)
    if err != nil {

        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
        return
    }

    if n != len(msg) {
        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    }

}