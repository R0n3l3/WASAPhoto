package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes
	rt.router.POST("/session/", rt.login)

	rt.router.GET("/users/", rt.getUserProfile)

	rt.router.PUT("/users/:userId", rt.setMyUserName)

	rt.router.POST("/users/:userId/banned/", rt.banUser)
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.unbanUser)

	rt.router.POST("/users/:userId/profile/photos/", rt.uploadPhoto)
	rt.router.DELETE("/users/:userId/profile/photos/:photoId", rt.deletePhoto)

	rt.router.POST("/users/:userId/profile/photos/:photoId/likes/", rt.likePhoto)
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/likes/:likeId", rt.unlikePhoto)

	rt.router.POST("/users/:userId/profile/photos/:photoId/comments/", rt.commentPhoto)
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/comments/:commentId", rt.uncommentPhoto)

	rt.router.GET("/users/:userId/profile/following/", rt.getMyStream)
	rt.router.POST("/users/:userId/profile/following/", rt.followUser)
	rt.router.DELETE("/users/:userId/profile/following/:followingId", rt.unfollowUser)

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
