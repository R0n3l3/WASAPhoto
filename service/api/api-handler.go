package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//  Register routes
	rt.router.POST("/session/", rt.doLogin)

	rt.router.GET("/users/:username/profile/", rt.getMyStream)

	rt.router.GET("/users/:username", rt.getUserProfile)
	rt.router.PUT("/users/:username", rt.setMyUsername)

	rt.router.POST("/users/:username/banned/", rt.banUser)
	rt.router.DELETE("/users/:username/banned/:bannedId", rt.unbanUser)

	rt.router.GET("/users/:username/profile/photos/", rt.getMyPhotos)
	rt.router.POST("/users/:username/profile/photos/", rt.uploadPhoto)
	rt.router.DELETE("/users/:username/profile/photos/:photoId", rt.deletePhoto)

	rt.router.GET("/users/:username/profile/photos/:photoId/likes/", rt.getLikes)
	rt.router.POST("/users/:username/profile/photos/:photoId/likes/", rt.likePhoto)
	rt.router.DELETE("/users/:username/profile/photos/:photoId/likes/:likeId", rt.unlikePhoto)

	rt.router.GET("/users/:username/profile/photos/:photoId/comments/", rt.getComments)
	rt.router.POST("/users/:username/profile/photos/:photoId/comments/", rt.commentPhoto)
	rt.router.DELETE("/users/:username/profile/photos/:photoId/comments/:commentId", rt.uncommentPhoto)

	rt.router.GET("/users/:username/profile/following/", rt.getFollowers)
	rt.router.POST("/users/:username/profile/following/", rt.followUser)
	rt.router.DELETE("/users/:username/profile/following/:followingId", rt.unfollowUser)

	return rt.router
}
