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
	rt.router.POST("/users/", rt.createUser)
	rt.router.GET("/users/:userId", rt.getUserProfile)
	rt.router.POST("/users/:userId/ban", rt.banUser)
	rt.router.GET("/users/:userId/stream", rt.getMyStream)
	rt.router.PUT("/users/:userId/name", rt.setMyUserName)
	rt.router.POST("/users/:userId/unban", rt.unbanUser)
	rt.router.POST("/users/:userId/follow", rt.followUser)
	rt.router.POST("/users/:userId/unfollow", rt.unfollowUser)
	rt.router.POST("/photos/:photoId/comments/", rt.commentPhoto)
	rt.router.DELETE("/comments/:commentId", rt.uncommentPhoto)
	rt.router.POST("/photos/:photoId/likes/", rt.likePhoto)
	rt.router.DELETE("/likes/:likeId", rt.unlikePhoto)
	rt.router.POST("/users/:userId/photos/", rt.uploadPhoto)
	rt.router.DELETE("/photos/:photoId", rt.deletePhoto)

	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
