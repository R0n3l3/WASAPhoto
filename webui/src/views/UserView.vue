
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
    newName:"",
    username: localStorage.getItem("searchName"),
    token: localStorage.getItem("token"),
    photoNumber: parseInt(localStorage.getItem("searchPhoto")),
    profile: {
      photoNumber: 0,
      profileId: 0,
      profileName: "",
    },
    comment: {
      commentId: 0,
      commenter: "",
      commentTime: "",
      content : "",
      photoComment: 0,
    },
    photo: {
      photoId: 0,
      uploader: 0,
      uploadTime: "",
      likeNumber: 0,
      commentNumber: 0,
      image: [],
    },
    myPhotos: [],
    photoData: [],
    comments: [],
    likes: [],
    followers: [],
    following: [],
    upload: null,
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
    } catch (e) {
      if (e.response && e.response.status !== 404) {
        this.errormsg = e.toString()
      }
    }
    this.photoData.push({
      photoId: photo.PhotoId,
      photoComments: this.comments,
    })
  },

  async refresh() {
    try {
      let response = await this.$axios.get("/users/" + this.username + "/profile/following/", {
        params: {
          followers: "true"
        }, headers:
            {
              Authorization: "Bearer " + this.token
            }
      })
      if (response.data !== null) {
        this.followers = response.data
      } else {
        this.followers = []
        this.followmsg = "You don't have followers"
      }
    } catch (e) {
      this.errormsg = e.toString()
    }
    try {
      let response = await this.$axios.get("/users/" + this.username + "/profile/following/", {
        params: {
          followers: "false"
        }, headers:
            {
              Authorization: "Bearer " + this.token
            }
      })
      if (response.data !== null) {
        this.following = response.data
      } else {
        this.following = []
        this.followingmsg = "You don't follow anyone"
      }
    } catch (e) {
      this.errormsg = e.toString()
    }

    try {
      let response = await this.$axios.get("/users/" + this.username + "/profile/photos/", {
        headers:
            {
              Authorization: "Bearer " + this.token
            }
      })
      if (response.data !== null) {
        this.myPhotos = response.data
      } else {
        this.myPhotos = []
        this.detailedmsg = "You haven't uploaded photos"
      }
      for (let i = 0; i < this.myPhotos.length; i++) {
        await this.getData(this.myPhotos[i])
      }
    } catch (e) {
      this.errormsg = e.toString()
    }

  },

  async changeUsername() {
    if (this.newName === "") {
      this.errormsg = "You have to type a new name"
    } else {
      this.errormsg = ""
      try {
        await this.$axios.put("/users/" + this.username, this.newName, {
          headers: {
            Authorization: "Bearer " + this.token
          }
        })
        this.username = this.newName
        localStorage.setItem("username", this.username)
        this.newName = ""
        localStorage.setItem("searchName", this.username)
        this.refresh()
      } catch (e) {
        if (e.response && e.response.status === 409) {
          this.errormsg = "Username already taken"
        } else {
          this.errormsg = e.toString()
        }
      }
    }
  },

  async handleFile() {
    this.upload = this.$refs.file.files[0]
  },

  async uploadPhoto() {
    this.detailedmsg = ""
    if (this.upload === null) {
      this.errormsg = "Select a photo"
    } else {
      this.errormsg = ""
      try {
        let response = await this.$axios.post("/users/" + this.username + "/profile/photos/", this.upload, {
          headers: {
            Authorization: "Bearer " + this.token
          }
        })
        if (response.data !== null) {
          this.photo = response.data
          this.photoNumber = this.photoNumber + 1
          this.myPhotos.unshift(this.photo)
          localStorage.setItem("searchPhoto", this.photoNumber)
          this.getData(this.photo)
        }
      } catch (e) {
        this.errormsg = e.toString()
      }
    }
  },

  async deletePhoto(id) {
    try {
      await this.$axios.delete("/users/" + this.username + "/profile/photos/" + id, {
        headers:
            {
              Authorization: "Bearer " + this.token
            }
      })
      this.photoNumber = this.photoNumber - 1
      localStorage.setItem("searchPhoto", this.photoNumber)
      this.photoData = []
      this.refresh()
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
    if (this.commentContent === "") {
      this.errormsg = "Please type a comment"
    } else {
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
    if (comment.Commenter !== this.username) {
      this.errormsg = "You cannot delete this comment"
    } else {
      this.errormsg = ""
      try {
        await this.$axios.delete("/users/" + photo.Uploader + "/profile/photos/" + photo.PhotoId + "/comments/" + comment.CommentId, {
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
      <input type="text" v-model="newName" placeholder="Insert new username" class="form-control-sm">
      <button class="btn btn-outline-dark" @click="changeUsername">Change</button>
      <h1 class="center-horizontal h2">{{this.username}}'s Profile</h1>
      <button class="btn btn-outline-dark" @click="homepage">Homepage</button>
    </div>
    <div class="flex-md-nowrap align-items-center pt-3 pb-2 mb-3">
      <div class="center-vertical d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <p class="center-horizontal">You have uploaded {{this.photoNumber}} photo. To upload a new photo, </p>
      <input type="file" accept="image/*" @change="handleFile" ref="file">
        <button class="btn btn-outline-dark" @click="uploadPhoto">Upload</button>
      </div>
    </div>
    <ErrorMsg v-if="detailedmsg" :msg="detailedmsg" class="center-vertical center-horizontal"></ErrorMsg>
    <div class="border-bottom pt-3 pb-2 mb-3" v-for="photo in myPhotos" :key="photo.PhotoId">
      <button class="btn btn-outline-danger" @click="deletePhoto(photo.PhotoId)">Delete Photo</button>
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
      <div class="d-flex justify-content-between align-items-center">
        <p>This photo has {{photo.LikeNumber}} likes and {{photo.CommentNumber}} comments.</p>
      </div>
        <input type="text" v-model="commentContent" placeholder="Type a new comment" class="form-control-plaintext">
        <button class="btn btn-outline-success" @click="commentPhoto(photo)">Add a comment</button>
    </div>

    <p> Your followers: </p>
    <div class="pb-2 mb-3 with-border-color scroll-container-follow">
      <ErrorMsg v-if="followmsg" :msg="followmsg" class="center-vertical center-horizontal"></ErrorMsg>
      <div class=" pb-2 mb-3 d-flex" v-for="user in followers" :key="user.ProfileId">
        <button class="btn btn-outline-dark" style="margin-left: 10px" @click="goToProfile(user.ProfileId)"> {{user.ProfileName}}</button>
      </div>
      </div>
    <p>You follow: </p>
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

