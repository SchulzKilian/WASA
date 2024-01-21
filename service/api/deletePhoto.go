package api

import (
    "net/http"
    "strconv"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    photo_id := ps.ByName("photoId")
    ctx.Logger.Info(photo_id)
    db := ctx.Database
    photoid, err := strconv.Atoi(photo_id)
    if err != nil {
        // If an error occurs, send a 404 Not Found error
        http.Error(w, "Type mismatch", http.StatusBadRequest)
        return
    }
    err = db.DeletePhoto(photoid)
    if err != nil {
        // If an error occurs, send a 404 Not Found error
        http.Error(w, "Photo not found", http.StatusNotFound)
        return
    }

    // If no error, send a success response
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Photo deleted successfully"))

}