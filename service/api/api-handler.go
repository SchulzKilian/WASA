package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter" // Import the router package you are using
	"github.com/sirupsen/logrus"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

)

type _router struct {
    router     *httprouter.Router
    baseLogger logrus.FieldLogger  // This should match the type in Config
    db         database.AppDatabase 
}

// Assume you have a function to create a new _router instance
func NewRouter() *_router {
	return &_router{
		router: httprouter.New(),
		// Initialize other fields if needed
	}
}

// ... Your existing code ...

// Handler returns an instance of httprouter.Router that handles APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session", rt.wrap(doLogin))  // works
	rt.router.POST("/users/", rt.wrap(createUser))   // works
	rt.router.GET("/users/:name", rt.wrap(getUserProfile))   // works
	rt.router.POST("/users/:name/banned/", rt.wrap(banUser)) // works
	rt.router.GET("/stream", rt.wrap(getMyStream)) // works
	rt.router.PATCH("/users/:name", rt.wrap(setMyUserName)) // works
	rt.router.DELETE("/users/:name/banned/", rt.wrap(unbanUser)) //works
	rt.router.POST("/users/:name/followers/", rt.wrap(followUser))  // works
	rt.router.DELETE("/users/:name/followers/", rt.wrap(unfollowUser)) // works
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(commentPhoto)) // works
	rt.router.DELETE("/comments/:commentId", rt.wrap(uncommentPhoto)) // works
	rt.router.POST("/photos/:photoId/likes/", rt.wrap(likePhoto)) // works
	rt.router.DELETE("/likes/:likeId", rt.wrap(unlikePhoto)) // works
	rt.router.POST("/photos/", rt.wrap(uploadPhoto))   // works
	rt.router.DELETE("/photos/:photoId", rt.wrap(deletePhoto))  // works
	rt.router.GET("/liveness", rt.wrap(rt.liveness))

	return rt.router
}
