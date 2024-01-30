package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//  Register routes
	rt.router.POST("/session/", rt.doLogin)

	rt.router.GET("/users/:userId", rt.getUserProfile)
	rt.router.PUT("/users/:userId", rt.setMyUsername)

	rt.router.POST("/users/:userId/banned/", rt.banUser)
	rt.router.DELETE("/users/:userId/banned/:bannedId", rt.unbanUser)

	rt.router.POST("/users/:userId/profile/photos/", rt.uploadPhoto) // todo: check api
	rt.router.DELETE("/users/:userId/profile/photos/:photoId", rt.deletePhoto)

	rt.router.POST("/users/:userId/profile/photos/:photoId/likes/", rt.likePhoto)            // todo: check api
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/likes/:likeId", rt.unlikePhoto) // todo: check api

	rt.router.POST("/users/:userId/profile/photos/:photoId/comments/", rt.commentPhoto)
	rt.router.DELETE("/users/:userId/profile/photos/:photoId/comments/:commentId", rt.uncommentPhoto)

	rt.router.POST("/users/:userId/profile/following/", rt.followUser)                 //todo: handle banned users
	rt.router.DELETE("/users/:userId/profile/following/:followingId", rt.unfollowUser) //todo: check api

	rt.router.GET("/users/:userId/profile/following/", rt.getMyStream)
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
