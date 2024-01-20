package api

import (
    "fmt"
    "io"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database" // replace with your actual package import path
)

func doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
    // Placeholder logic
    bodyBytes, err := io.ReadAll(r.Body)
    if err != nil {
        http.Error(w, "Error reading request body", http.StatusInternalServerError)
        return
    }
    defer r.Body.Close()

    // Convert body bytes to string
    username := string(bodyBytes)
    ctx.Logger.Info(username)
    db := ctx.Database
    userexists, err, token := db.DoesUserExist(username)
    if err != nil {
        ctx.Logger.WithError(err).Error("Error checking user existence")
        http.Error(w, "Internal server error", http.StatusInternalServerError)
        return
    }
    ctx.Logger.Info("Logged user in") // Example logging
    if userexists{
        fmt.Fprintf(w, token)
        return

    }
    if !userexists{

        var user database.User
        
        user.Username = username
    
        // Call AddUser method of the database object
        ctx.Logger.Info("Made it to the adduser call")
    
        err,token = db.AddUser(&user)
    
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return 
        }
    
        ctx.Logger.Info("User created successfully")
        fmt.Fprintf(w, token)
        return
    }

}