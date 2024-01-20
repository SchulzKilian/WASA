package api

import (
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
    "encoding/json"
    "net/http"
    "github.com/julienschmidt/httprouter"
    "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/api/reqcontext"
	"fmt"
    // Import your database package here
)

// Assuming User struct is defined in your database package
// import your database package and refer to it as dbpkg

func createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

    // Decode the request body into a User struct
    db := ctx.Database
    var user database.User
    
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Call AddUser method of the database object
	ctx.Logger.Info("Made it to the adduser call")

    err, token := db.AddUser(&user)

    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return 
    }

    ctx.Logger.Info("User created successfully")
    fmt.Fprintf(w, token)
}
