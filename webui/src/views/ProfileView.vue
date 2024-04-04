<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: "",
			detailedmsg: "",
      followmsg: "",
      followingmsg:"",
      commentContent: "",
      token: localStorage.getItem("token"),
      username: localStorage.getItem("username"),
      profileName: localStorage.getItem("searchName"),
      photoNumber: parseInt(localStorage.getItem("searchPhoto")),
      profileId: localStorage.getItem("searchId"),
      photo: {
        photoId: 0,
        uploader: 0,
        uploadTime: "",
        likeNumber: 0,
        commentNumber: 0,
        image: [],
      },
      comment: {
        commentId: 0,
        commenter: "",
        commentTime: "",
        content : "",
        photoComment: 0,
      },
      like: {
        likeId: 0,
        liker: 0,
        photoLiked: 0,
      },
      photos: [],
      photoData: [],
      comments: [],
      likes: [],
      followers: [],
      following: [],
      upload:null,
    }
	},
	methods: {
		filteredPhotoData(photoId) {
			return this.photoData.filter(data => data.photoId === photoId);
		},

    async getData(photo) {
      try {
        let response = await this.$axios.get("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/comments/", {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data !== null) {
          this.comments = response.data
        } else {
          this.comments = []
        }
      }catch (e) {
        if (e.response && e.response.status !== 404) {
          this.errormsg = "Data not found"
        }else{
          this.errormsg = e.toString()
        }
      }
      try {
        let response = await this.$axios.get("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/likes/", {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data !== null) {
          this.likes = response.data
        } else {
          this.likes = []
        }
      }catch(e) {
        if (e.response && e.response.status !== 404) {
          this.errormsg = "Data not found"
        }else{
          this.errormsg = e.toString()
        }
      }
      this.photoData.push({
        photoId: photo.PhotoId,
        photoComments: this.comments,
        photoLikes: this.likes,
      })
    },

    async getFollow() {
      try {
        let response = await this.$axios.get("/users/"+this.profileName+"/profile/following/", {
          params: {
            followers: "true"
          }, headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data!==null) {
          this.followers=response.data
        }else{
          this.followers=[]
          this.followmsg="The user doesn't have followers"
        }
      }catch (e) {
        this.errormsg = e.toString()
      }
      try {
        let response = await this.$axios.get("/users/"+this.profileName+"/profile/following/", {
          params: {
            followers: "false"
          }, headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data!==null) {
          this.following=response.data
        }else{
          this.following=[]
          this.followingmsg="The user doesn't follow anyone"
        }
      }catch (e) {
        this.errormsg = e.toString()
      }

    },

    async refresh() {
      try {
        await this.$axios.get("/users/" + this.profileName + "/banned/", {
          params: {
            banned: this.username
          }, headers: {
            Authorization: "Bearer " + this.token
          }
        })
      }catch(e){
        this.$router.push({path: "/err"})
        }
      await this.getFollow()
      try {
        let response = await this.$axios.get("/users/" + this.profileName + "/profile/photos/", {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data!==null) {
          this.photos=response.data
        }else{
          this.photos=[]
          this.detailedmsg="The user hasn't uploaded photos"
        }
        for (let i=0; i<this.photos.length; i++) {
          await this.getData(this.photos[i])
        }
      }catch(e){
        this.$router.push({path: "/err"})
      }

    },

    async follow() {
      this.followmsg=""
      try {
        let response = await this.$axios.post("/users/"+this.username+"/profile/following/", this.profileName, {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        let profile = {
          photoNumber: response.data.PhotoNumber,
          profileId: response.data.ProfileId,
          profileName: response.data.ProfileName,
        }
        this.followers.push(profile)
        this.getFollow()
      }catch(e) {
        if (e.response && e.response.status === 409) {
          this.errormsg = "You already follow this user"
        }else {
          this.errormsg = e.toString()
        }
      }
    },

    async unfollow() {
      this.errormsg=""
      try {
        await this.$axios.delete("/users/"+this.username+"/profile/following/"+this.profileName, {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        this.followers = []
        this.getFollow()
      } catch (e) {
        this.errormsg = e.toString()
      }
    },

		async likePhoto(photo) {
			if (photo.Uploader===this.username) {
				this.errormsg="You cannot like your own photo"
			}else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/users/" + this.profileName + "/profile/photos/" + photo.PhotoId + "/likes/", this.username, {
						headers:
							{
								Authorization: "Bearer " + this.token
							}
					})
					this.like=response.data
					this.photoData.forEach(p=>{
						if (p.photoId===photo.PhotoId) {
							p.photoLikes.push(this.like)
						}
					})
					photo.LikeNumber+=1
				} catch (e) {
          if (e.response && e.response.status === 409) {
            this.errormsg = "You already liked this photo"
				}else{
          this.errormsg=e.toString()}
        }
			}
		},

		async checkLikes(photo) {
      this.errormsg=""
      this.photoData.forEach(p => {
        if (p.photoId === photo.PhotoId) {
          this.likes = p.photoLikes
        }
        if (this.likes.length === 0) {
          this.errormsg = "You haven't liked this photo"
        } else {
					this.likes.forEach(l => {
						if (parseInt(l.Liker) === parseInt(this.token)) {
							this.unlikePhoto(photo, l)
						} else{
							this.errormsg = "You haven't liked this photo"
						}
					})
				}
			})
		},

		async unlikePhoto(photo, l) {
      this.errormsg=""
			try {
				await this.$axios.delete("/users/" + this.profileName + "/profile/photos/"+photo.PhotoId+"/likes/"+l.LikeId, {
					headers:
						{
							Authorization: "Bearer " + this.token
						}
				})
				this.photoData = []
				this.refresh()
			} catch (e) {
				this.errormsg = e.toString()
			}
		},

		async commentPhoto(photo) {
			if (this.commentContent==="") {
				this.errormsg="Please type a comment"
			}else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/users/"+this.profileName+"/profile/photos/"+photo.PhotoId+"/comments/", {
						content: this.commentContent,
						commenter: this.username
					}, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.comment=response.data
					this.photoData.forEach(p=>{
						if (p.photoId===photo.PhotoId) {
							p.photoComments.push(this.comment)
						}
					})
					photo.CommentNumber+=1
					this.content=""
				}catch(e){
					this.errormsg = e.toString()
				}
			}
		},

		async deleteComment(photo, comment) {
			if (comment.Commenter!==this.username) {
				this.errormsg="You cannot delete this comment"
			}else {
				this.errormsg=""
				try {
					await this.$axios.delete("/users/" + this.profileName + "/profile/photos/" + photo.PhotoId + "/comments/" + comment.CommentId, {
						headers:
							{
								Authorization: "Bearer " + this.token
							}
					})
					this.photoData = []
					this.refresh()
				} catch (e) {
					this.errormsg = e.toString()
				}
			}
		},

    async ban() {
      try {
        await this.$axios.post("/users/"+this.username+"/banned/", this.profileName, {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        this.errormsg="Ban successful"
        this.followers = []
        this.following=[]
        this.photoData=[]
        this.refresh()
      }catch(e) {
        if (e.response && e.response.status === 409) {
          this.errormsg = "You already banned this person"
        } else {
          this.errormsg = e.toString()
        }
      }
    },

    async unban() {
      try {
        await this.$axios.delete("/users/"+this.username+"/banned/"+this.profileName, {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
      }catch(e){
        this.errormsg = e.toString()
      }
    },
    async homepage() {
      localStorage.setItem("searchName", "")
      localStorage.setItem("searchId", "")
      localStorage.setItem("searchPhoto", "")
      this.$router.push({path: "/session/"})
    }
	},
	mounted() {
		this.refresh()

	}
}
</script>

<template>
  <div>
    <div class="center-vertical d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="center-horizontal h2">{{this.profileName}}'s Profile</h1>
      <button class="btn btn-outline-dark" @click="homepage">Homepage</button>
    </div>
    <div class="center-vertical d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <ErrorMsg v-if="detailedmsg" :msg="detailedmsg" class="center-vertical center-horizontal"></ErrorMsg>
      <button class="btn btn-outline-success" @click="follow">Follow User</button>
      <button class="btn btn-outline-danger" @click="unfollow">Unfollow User</button>
      <p class="center-horizontal">This user uploaded {{this.photoNumber}} photos. </p>
      <button class="btn btn-outline-success" @click="ban">Ban User</button>
      <button class="btn btn-outline-danger" @click="unban">Unban User</button>
    </div>
      <div class="border-bottom pt-3 pb-2 mb-3" v-for="photo in photos" :key="photo.PhotoId">
      <div class="border-bottom pt-3 pb-2 mb-3 d-flex flex-wrap flex-md-nowrap align-items-center">
        <img :src="'data:image/*; base64,' + photo.Image" alt="photo" class="resizable-image with-border">
        <div class="scroll-container with-border-color" v-for="data in filteredPhotoData(photo.PhotoId)" :key="data.PhotoId">
          <div v-for="comment in data.photoComments" :key="comment.CommentId" class="align-items-center justify-content-between flex-wrap" style="padding: 10px;">
            {{comment.Commenter}} commented: "{{comment.Content}}"
            <button class="btn btn-outline-danger" @click="deleteComment(photo, comment)">Delete Comment</button>
          </div>
        </div>
      </div>
      <p>This photo was uploaded on {{photo.UploadTime}}.</p>
      <p>This photo has {{photo.LikeNumber}} likes and {{photo.CommentNumber}} comments.</p>
      <div>
        <button class="btn btn-outline-success" @click="likePhoto(photo)">Add a like</button>
        <button class="btn btn-outline-danger" @click="checkLikes(photo)">Remove a like</button>
      </div>
      <input type="text" v-model="commentContent" placeholder="Type a new comment" class="form-control-plaintext">
      <button class="btn btn-outline-success" @click="commentPhoto(photo)">Add a comment</button>
    </div>
    <p> Followers: </p>
    <div class="pb-2 mb-3 with-border-color scroll-container-follow">
      <ErrorMsg v-if="followmsg" :msg="followmsg" class="center-vertical center-horizontal"></ErrorMsg>
      <div class=" pb-2 mb-3 d-flex" v-for="user in followers" :key="user.ProfileId">
        <button class="btn btn-outline-dark" style="margin-left: 10px" @click="goToProfile(user.ProfileId)"> {{user.ProfileName}}</button>
      </div>
    </div>
    <p>Following: </p>
    <div class="pt-3 mb-3 with-border-color scroll-container-follow">
      <ErrorMsg v-if="followingmsg" :msg="followingmsg" class="center-vertical center-horizontal"></ErrorMsg>
      <div class="pt-3 mb-3 d-flex" v-for="user in following" :key="user.ProfileId">
        <button class="btn btn-outline-dark" style="margin-left: 10px" @click="goToProfile(user.ProfileId)"> {{user.ProfileName}}</button>
      </div>
    </div>
    <div>
      <ErrorMsg v-if="errormsg" :msg="errormsg" class="center-vertical center-horizontal"></ErrorMsg>
    </div>
  </div>
</template>

<style>
</style>
