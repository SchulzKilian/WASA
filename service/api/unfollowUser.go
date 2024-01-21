package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    ctx.Logger.Info("myApiHandler called") // Example logging
    username := ctx.User.Username
    tofollow := ps.ByName("name")
    db := ctx.Database
    err := db.DeleteFollow(username,tofollow)
    if err !=nil{
        http.Error(w,"Error unfollowing the user",http.StatusBadRequest)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Successfully unfollowed the user"))
}