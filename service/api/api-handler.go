package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter" // Import the router package you are using
)

type _router struct {
	router *httprouter.Router
	// Add any other necessary fields
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
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.POST("/users/", rt.wrap(rt.createUser))
	rt.router.GET("/users/:userId", rt.wrap(rt.getUserProfile))
	rt.router.POST("/users/:userId/ban", rt.wrap(rt.banUser))
	rt.router.GET("/users/:userId/stream", rt.wrap(rt.getMyStream))
	rt.router.PUT("/users/:userId/name", rt.wrap(rt.setMyUserName))
	rt.router.POST("/users/:userId/unban", rt.wrap(rt.unbanUser))
	rt.router.POST("/users/:userId/follow", rt.wrap(rt.followUser))
	rt.router.POST("/users/:userId/unfollow", rt.wrap(rt.unfollowUser))
	rt.router.POST("/photos/:photoId/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/comments/:commentId", rt.wrap(rt.uncommentPhoto))
	rt.router.POST("/photos/:photoId/likes/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/likes/:likeId", rt.wrap(rt.unlikePhoto))
	rt.router.POST("/users/:userId/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:photoId", rt.wrap(rt.deletePhoto))

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
