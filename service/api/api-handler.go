package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	//  Register routes
	rt.router.POST("/session/", rt.wrap(rt.doLogin))

	rt.router.GET("/users/:username", rt.wrap(rt.getUserProfile))
	rt.router.PUT("/users/:username", rt.wrap(rt.setMyUsername))

	rt.router.GET("/users/:username/banned/", rt.wrap(rt.getBan))
	rt.router.POST("/users/:username/banned/", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:username/banned/:bannedId", rt.wrap(rt.unbanUser))

	rt.router.GET("/users/:username/profile/", rt.wrap(rt.getMyStream))

	rt.router.GET("/users/:username/profile/photos/", rt.wrap(rt.getMyPhotos))
	rt.router.POST("/users/:username/profile/photos/", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/users/:username/profile/photos/:photoId", rt.wrap(rt.deletePhoto))

	rt.router.GET("/users/:username/profile/photos/:photoId/likes/", rt.wrap(rt.getLikes))
	rt.router.POST("/users/:username/profile/photos/:photoId/likes/", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/users/:username/profile/photos/:photoId/likes/:likeId", rt.wrap(rt.unlikePhoto))

	rt.router.GET("/users/:username/profile/photos/:photoId/comments/", rt.wrap(rt.getComments))
	rt.router.POST("/users/:username/profile/photos/:photoId/comments/", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/users/:username/profile/photos/:photoId/comments/:commentId", rt.wrap(rt.uncommentPhoto))

	rt.router.GET("/users/:username/profile/following/", rt.wrap(rt.getFollowers))
	rt.router.POST("/users/:username/profile/following/", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:username/profile/following/:followingId", rt.wrap(rt.unfollowUser))

	return rt.router
}
