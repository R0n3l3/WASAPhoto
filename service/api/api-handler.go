package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session/", rt.doLogin) //todo: errors

	rt.router.GET("/users/", rt.getUserProfile) //todo: errors

	rt.router.PUT("/users/:userId", rt.setMyUserName) //todo: errors

	rt.router.POST("/users/:userId/banned/", rt.banUser)              //todo: errors
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.unbanUser) //todo: errors

	rt.router.POST("/users/:userId/profile/photos/", rt.uploadPhoto)           //todo: errors
	rt.router.DELETE("/users/:userId/profile/photos/:photoId", rt.deletePhoto) //todo: errors

	rt.router.POST("/users/:userId/profile/photos/:photoId/likes/", rt.likePhoto)            //todo
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/likes/:likeId", rt.unlikePhoto) //todo

	rt.router.POST("/users/:userId/profile/photos/:photoId/comments/", rt.commentPhoto)               //todo
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/comments/:commentId", rt.uncommentPhoto) //todo

	rt.router.GET("/users/:userId/profile/following/", rt.getMyStream)                 //todo
	rt.router.POST("/users/:userId/profile/following/", rt.followUser)                 //todo: errors
	rt.router.DELETE("/users/:userId/profile/following/:followingId", rt.unfollowUser) //todo: errors

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
