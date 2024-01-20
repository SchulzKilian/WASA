package api

import (
    "net/http"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    name := ps.ByName("name")
    db := ctx.Database
    ctx.Logger.Info("Right before get User details")
    details, err := db.GetUserDetails(name)
    if err != nil{
        http.Error(w,err.Error(),http.StatusBadRequest)
        return
    }
    ctx.Logger.Info("Right after get User details")

    jsonResponse, err := json.Marshal(details)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    ctx.Logger.Info("myApiHandler called") // Example logging

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)

}