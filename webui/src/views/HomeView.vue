<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			loading: false,
			search: "",
			username: localStorage.getItem("username"),
			token: localStorage.getItem("token"),
			photoStream: [
			],
			profile: {
				profileId: 0,
				profileName: "",
				photoNumber: 0,
			},
			comments: [],
			comment: {
				commentId: 0,
				commenter: 0,
				commentTime: "",
				content: "",
				photoComment: 0,
			},
			photosComments: [],
			commentContent: "",
      photoData: [
      ],
      likes: [],
      like: {
        likeId: 0,
        liker: 0,
        photoLiked: 0,
      },
		}
	},
	methods: {
    filteredPhotoData(photoId) {
      return this.photoData.filter(data => data.photoId === photoId);
    },
    async getData(photo) {
      let response = await this.$axios.get("/users/" + this.profileName + "/profile/photos/" + photo.PhotoId + "/comments/", {
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

      response = await this.$axios.get("/users/" + this.profileName + "/profile/photos/" + photo.PhotoId + "/likes/", {
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
          this.photoStream = response.data
          for (let i = 0; i < this.photoStream.length; i++) {
            await this.getData(this.photoStream[i])
          }
        } else {
          this.photoStream = []
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
        this.errormsg = null
        try {
          let response = await this.$axios.get("/users/" + this.username, {
            params: {
              search: this.search
            },
            headers: {
              Authorization: "Bearer " + this.token
            }
          })
          let decodedprofile = response.data
          if (decodedprofile.ProfileName === this.username) {
            this.errormsg = "You cannot search yourself! Please use the 'My profile' button"
          } else {
            this.profile.profileName = decodedprofile.ProfileName
            this.profile.profileId = decodedprofile.ProfileId
            this.profile.photoNumber = decodedprofile.PhotoNumber
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

    async getMyProfile() {
      try {
        let response = await this.$axios.get("/users/" + this.username, {
          params: {
            search: this.username
          },
          headers: {
            Authorization: "Bearer " + this.token
          }
        })
        let decodedprofile = response.data
        this.profile.profileName = decodedprofile.ProfileName
        this.profile.profileId = decodedprofile.ProfileId
        this.profile.photoNumber = decodedprofile.PhotoNumber
        localStorage.setItem("searchName", this.profile.profileName)
        localStorage.setItem("searchId", this.profile.profileId)
        localStorage.setItem("searchPhoto", this.profile.photoNumber)
        this.$router.push({path: "/users/" + this.profile.profileName + "/view/"})

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
          this.commentContent=""
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
          this.getStream()
        } catch (e) {
          this.errormsg = e.toString()
        }
      }
    },

    async likePhoto(photo) {
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
          this.errormsg = e.toString()
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
        this.getStream()
      } catch (e) {
        this.errormsg = e.toString()
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
        let decodedprofile = response.data
        this.profile.profileName = decodedprofile.ProfileName
        this.profile.profileId = decodedprofile.ProfileId
        this.profile.photoNumber = decodedprofile.PhotoNumber
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
  },



	mounted() {
		this.getStream()
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{this.username}} Homepage</h1>
		</div>
		<div>
			<input type="text" v-model="search" placeholder="Insert a user to search"/>
			<button @click="getUser">Search</button>
		</div>
		<div>
			<button @click="getMyProfile">Your Profile</button>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p> Your stream</p>
      <div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="photo in photoStream" :key="photo.PhotoId">
        <button @click="goToProfile(photo.Uploader)"> Go to uploader page</button>
        <img :src="'data:image/*; base64,' + photo.Image" alt="photo" class="resizable-image">
        <p> Comment number: {{photo.CommentNumber}}</p>
        <div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2" v-for="data in filteredPhotoData(photo.PhotoId)" :key="data.PhotoId">
          <div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="comment in data.photoComments" :key="comment.CommentId">
            {{comment.Commenter}} --> {{comment.Content}}
            <button @click="deleteComment(photo, comment)">Delete Comment</button>
          </div>
        </div>
        <p>Like number: {{photo.LikeNumber}}</p>
        <button @click="likePhoto(photo)">Add a like</button>
        <button @click="checkLikes(photo)">Remove a like</button>
        <input type="text" v-model="commentContent" placeholder="Type a comment">
        <button @click="commentPhoto(photo)">Add a comment</button>
      </div>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
