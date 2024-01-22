package api

import (
    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext" // replace with your actual package import path
)

func followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    if ctx.User == nil{
        http.Error(w, "You have to be logged in to like", http.StatusForbidden)
        return
    }
    ctx.Logger.Info("myApiHandler called") // Example logging
    username := ctx.User.Username
    tofollow := ps.ByName("name")
    db := ctx.Database
    err := db.AddFollow(username,tofollow)
    if err !=nil{
        http.Error(w,"Error following the user",http.StatusBadRequest)
        ctx.Logger.Info(err)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Successfully followed the user"))
}