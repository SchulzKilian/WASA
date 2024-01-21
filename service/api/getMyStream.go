package api

import (
    "net/http"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    //"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    db := ctx.Database
    tooutput, err := db.GetFollowedUsersPhotos(ctx.User.Username)
    if err != nil{
        http.Error(w,err.Error(),http.StatusBadRequest)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // Serialize the tooutput slice to JSON
    jsonResponse, err := json.Marshal(tooutput)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Write the JSON response
    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)    
}