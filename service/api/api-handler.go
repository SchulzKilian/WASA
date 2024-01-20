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
	rt.router.POST("/session", rt.wrap(doLogin))
	rt.router.POST("/users/", rt.wrap(createUser))
	rt.router.GET("/users/:userId", rt.wrap(getUserProfile))
	rt.router.POST("/users/:userId/ban", rt.wrap(banUser))
	rt.router.GET("/users/:userId/stream", rt.wrap(getMyStream))
	rt.router.PATCH("/users/:name", rt.wrap(setMyUserName))
	rt.router.POST("/users/:userId/unban", rt.wrap(unbanUser))
	rt.router.POST("/users/:userId/follow", rt.wrap(followUser))
	rt.router.POST("/users/:userId/unfollow", rt.wrap(unfollowUser))
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(commentPhoto))
	rt.router.DELETE("/comments/:commentId", rt.wrap(uncommentPhoto))
	rt.router.POST("/photos/:photoId/likes/", rt.wrap(likePhoto))
	rt.router.DELETE("/likes/:likeId", rt.wrap(unlikePhoto))
	rt.router.POST("/users/:userId/photos/", rt.wrap(uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(deletePhoto))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
