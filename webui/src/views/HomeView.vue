<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: "",
      detailedmsg: "",
      search: "",
      commentContent: "",
      username: localStorage.getItem("username"),
      token: localStorage.getItem("token"),
      profile: {
        profileId: 0,
        profileName: "",
        photoNumber: 0,
      },
      comment: {
        commentId: 0,
        commenter: 0,
        commentTime: "",
        content: "",
        photoComment: 0,
      },
      like: {
        likeId: 0,
        liker: 0,
        photoLiked: 0,
      },
      stream: [],
      comments: [],
      photosComments: [],
      photoData: [],
      likes: [],
		}
	},
	methods: {
    filteredPhotoData(photoId) {
      return this.photoData.filter(data => data.photoId === photoId);
      },

    filteredLikes(userId) {
      return this.likes.filter(l => parseInt(l.Liker) === userId);
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
          this.errormsg = e.toString()
        }
      }
      this.photoData.push({
        photoId: photo.PhotoId,
        photoComments: this.comments,
        photoLikes: this.likes,
      })
    },

    async getStream() {
      try {
        let response = await this.$axios.get("/users/" + this.username + "/profile/", {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        if (response.data !== null) {
          this.stream = response.data
          for (let i = 0; i < this.stream.length; i++) {
            await this.getData(this.stream[i])
          }
        } else {
          this.stream = []
          this.detailedmsg = "Your stream is empty"
        }
      } catch (e) {
        this.errormsg = e.toString()
      }

    },

    async getUser() {
      if (this.search === "") {
        this.errormsg = "You must insert a username"
      } else {
        this.errormsg = ""
        try {
          let response = await this.$axios.get("/users/" + this.username, {
            params: {
              search: this.search
            },
            headers: {
              Authorization: "Bearer " + this.token
            }
          })
          this.profile.profileName = response.data.ProfileName
          this.profile.profileId = response.data.ProfileId
          this.profile.photoNumber = response.data.PhotoNumber
          if (this.profile.profileName === this.username) {
            this.errormsg = "You cannot search yourself! Please use the 'My profile' button in the nav bar"
          } else {
            localStorage.setItem("searchName", this.profile.profileName)
            localStorage.setItem("searchId", this.profile.profileId)
            localStorage.setItem("searchPhoto", this.profile.photoNumber)
            this.$router.push({path: "/users/" + this.profile.profileName + "/profile"})
          }
        } catch (e) {
          if (e.response && e.response.status === 404) {
            this.errormsg = "User not found"
          } else {
            this.errormsg = e.toString()
          }

        }
      }
    },

    async goToProfile(id) {
      try {
        let response = await this.$axios.get("/users/" + this.username, {
          params: {
            search: id
          },
          headers: {
            Authorization: "Bearer " + this.token
          }
        })
        this.profile.profileName = response.data.ProfileName
        this.profile.profileId = response.data.ProfileId
        this.profile.photoNumber = response.data.PhotoNumber
        localStorage.setItem("searchName", this.profile.profileName)
        localStorage.setItem("searchId", this.profile.profileId)
        localStorage.setItem("searchPhoto", this.profile.photoNumber)
        this.$router.push({path: "/users/" + this.profile.profileName + "/profile"})
      } catch (e) {
        if (e.response && e.response.status === 404) {
          this.errormsg = "User not found"
        } else {
          this.errormsg = e.toString()
        }
      }

    },

    async commentPhoto(photo) {
      if (this.commentContent==="") {
        this.errormsg="Please type a comment"
      }else {
        this.errormsg = ""
        try {
          let response = await this.$axios.post("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/comments/", {
            content: this.commentContent,
            commenter: this.username
          }, {
            headers: {
              Authorization: "Bearer " + this.token
            }
          })
          this.comment = response.data
          let p = this.filteredPhotoData(photo.PhotoId)
          p[0].photoComments.push(this.comment)
          photo.CommentNumber += 1
          this.commentContent = ""
        } catch (e) {
          if (e.response && e.response.status === 404) {
            this.errormsg = "Not found"
          } else {
            this.errormsg = e.toString()
          }
        }
      }
    },

    async deleteComment(photo, comment) {
      if (comment.Commenter!==this.username) {
        this.errormsg="You cannot delete this comment"
      }else {
        this.errormsg=""
        try {
          await this.$axios.delete("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/comments/" + comment.CommentId, {
            headers:
                {
                  Authorization: "Bearer " + this.token
                }
          })
          this.photoData = []
          this.getStream()
        } catch (e) {
          this.errormsg = e.toString()
        }
      }
    },

    async checkLikes(photo) {
      this.errormsg=""
      let p = this.filteredPhotoData(photo.PhotoId)
      this.likes = p[0].photoLikes
      if (this.likes.length === 0) {
        this.errormsg = "You haven't liked this photo"
      } else {
        let l = this.filteredLikes(parseInt(this.token))
        if (l!==null){
          this.unlikePhoto(photo, l[0])
            } else{
              this.errormsg = "You haven't liked this photo"
            }
          }
        },

    async likePhoto(photo) {
      this.errormsg=""
        try {
          let response = await this.$axios.post("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/likes/", this.username, {
            headers:
                {
                  Authorization: "Bearer " + this.token
                }
          })
          this.like=response.data
          let p = this.filteredPhotoData(photo.PhotoId)
          p[0].photoLikes.push(this.like)
          photo.LikeNumber+=1
        } catch (e) {
          if (e.response && e.response.status === 409) {
            this.errormsg = "You already liked this photo"
          }else {
            this.errormsg = e.toString()
          }
        }
    },

    async unlikePhoto(photo, l) {
      this.errormsg=""
      try {
        await this.$axios.delete("/users/" + photo.Uploader + "/profile/photos/"+photo.PhotoId+"/likes/"+l.LikeId, {
          headers:
              {
                Authorization: "Bearer " + this.token
              }
        })
        this.photoData = []
        this.getStream()
      } catch (e) {
        this.errormsg = e.toString()
      }
    },
  },
	mounted() {
		this.getStream()
	}
}
</script>

<template>
	<div>
		<div
			class="center-vertical d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <input type="text" v-model="search" placeholder="Insert a user to search" class="form-control-sm">
      <button class="btn btn-outline-dark" @click="getUser">Search</button>
      <h1 class="center-horizontal h2">{{this.username}}'s Homepage</h1>
    </div>
		<div class="border-bottom pt-3 pb-2 mb-3" v-for="photo in stream" :key="photo.PhotoId">
      <button class="btn btn-outline-dark" @click="goToProfile(photo.Uploader)"> Go to uploader page</button>
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
		<ErrorMsg v-if="errormsg" :msg="errormsg" class="center-vertical center-horizontal"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg" class="center-vertical center-horizontal"></ErrorMsg>
	</div>
</template>

<style>
</style>
