package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    ctx.Logger.Info("myApiHandler called") // Example logging
    username := ctx.User.Username
    toban := ps.ByName("name")
    db := ctx.Database
    err := db.AddBan(username,toban)
    if err !=nil{
        http.Error(w,"Error banning the user",http.StatusBadRequest)
        ctx.Logger.Info(err)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Successfully banned the user"))
}