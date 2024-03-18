<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: "",
			detailedmsg: "",
			token: localStorage.getItem("token"),
			username: localStorage.getItem("searchName"),
			photoNumber: parseInt(localStorage.getItem("searchPhoto")),
			myPhotos: [],
			photoData: [
			],
			comment: {
				commentId: 0,
				commenter: "",
				commentTime: "",
				content : "",
				photoComment: 0,
			},
			comments: [],
			content: "",
			likes: [],
			upload:null,
			photo: {
				photoId: 0,
				uploader: 0,
				uploadTime: "",
				likeNumber: 0,
				commentNumber: 0,
				image: [],
			},
			newName:"",
      profile: {
        photoNumber: 0,
        profileId: 0,
        profileName: "",
      },
      followers: [],
      following: [],

		}
	},
	methods: {
		filteredPhotoData(photoId) {
			// Filter photoData based on photoId
			return this.photoData.filter(data => data.photoId === photoId);
		},

		async getData(photo) {
			let response = await this.$axios.get("/users/"+this.username+"/profile/photos/"+photo.PhotoId+"/comments/", {
				headers:
					{
						Authorization: "Bearer " + this.token
					}
			})
			if (response.data!==null) {
				this.comments = response.data
			}else {
				this.comments=[]
			}
			this.photoData.push({
				photoId: photo.PhotoId,
				photoComments: this.comments,
				photoLikes: this.likes,
			})
		},

		async refresh() {
      try {
        let response = await this.$axios.get("/users/"+this.username+"/profile/following/", {
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
          this.detailedmsg="You don't have followers"
        }
      }catch (e) {
        this.errormsg = e.toString()
      }
      //Get following
      try {
        let response = await this.$axios.get("/users/"+this.username+"/profile/following/", {
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
          this.errormsg="You don't follow anyone"
        }
      }catch (e) {
        this.errormsg = e.toString()
      }
			//Get my photos
			try {
				let response = await this.$axios.get("/users/" + this.username + "/profile/photos/", {
					headers:
						{
							Authorization: "Bearer " + this.token
						}
				})
				if (response.data!==null) {
					this.myPhotos=response.data
				}else{
					this.myPhotos=[]
					this.detailedmsg="You haven't uploaded photos"
				}
				for (let i=0; i<this.myPhotos.length; i++) {
					await this.getData(this.myPhotos[i])
				}
			}catch(e){
				this.errormsg = e.toString()
			}

		},

		async changeUsername(){
			if (this.newName==="") {
				this.errormsg="You have to type a new name"
			}else{
				this.errormsg=""
				try {
					await this.$axios.put("/users/"+this.username, this.newName, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.username=this.newName
					this.newName=""
					localStorage.setItem("searchName", this.username)
					localStorage.setItem("username", this.username)
					this.refresh()
				}catch(e) {
					this.errormsg = e.toString()
				}
			}
		},

		async handleFile() {
			this.upload=this.$refs.file.files[0]
		},

		async uploadPhoto() {
			this.detailedmsg=""
			if (this.upload===null) {
				this.errormsg="Select a photo"
			}else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/users/"+this.username+"/profile/photos/", this.upload, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					if (response.data!==null) {
					this.photo=response.data}
					this.photoNumber=this.photoNumber+1
					this.myPhotos.unshift(this.photo)
					localStorage.setItem("searchPhoto", this.photoNumber)
					this.getData(this.photo)
				}catch(e) {
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
				this.photoNumber=this.photoNumber-1
				localStorage.setItem("searchPhoto", this.photoNumber)
				this.photoData=[]
				this.refresh()
			}catch(e) {
				this.errormsg = e.toString()
			}
		},

		async commentPhoto(photo) {
			if (this.content==="") {
				this.errormsg="Please type a comment"
			}else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/users/"+this.username+"/profile/photos/"+photo.PhotoId+"/comments/", {
						content:this.content,
						commenter: this.username,
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
			try {
				await this.$axios.delete("/users/"+this.username+"/profile/photos/"+photo.PhotoId+"/comments/"+comment.CommentId, {
					headers:
						{
							Authorization: "Bearer " + this.token
						}
				})
				photo.CommentNumber-=1
				this.photoData=[]
				this.refresh()
			}catch(e) {
				this.errormsg = e.toString()
			}
			}
		},

		async goBack() {
			localStorage.setItem("searchName", "")
			localStorage.setItem("searchId", "")
			localStorage.setItem("searchPhoto", "")
			this.$router.push({path: "/session/"})
		},
	},
		mounted() {
			this.refresh()

	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{this.username}} Profile</h1>
			You uploaded {{this.photoNumber}} photos
		</div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<input type="text" v-model="newName" placeholder="Write a new name">
			<button @click="changeUsername">Change Username</button>
		</div>
		<div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="photo in myPhotos" :key="photo.PhotoId">
			<img :src="'data:image/*; base64,' + photo.Image" alt="photo" class="resizable-image">
			<p> Comment number: {{photo.CommentNumber}}</p>
			<div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2" v-for="data in filteredPhotoData(photo.PhotoId)" :key="data.PhotoId">
				<div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="comment in data.photoComments" :key="comment.CommentId">
					{{comment.Commenter}} --> {{comment.Content}}
					<button @click="deleteComment(photo, comment)">Delete Comment</button>
				</div>
			</div>
			<p>Like number: {{photo.LikeNumber}}</p>
			<input type="text" v-model="content" placeholder="Type a comment">
			<button @click="commentPhoto(photo)">Add a comment</button>
			<button @click="deletePhoto(photo.PhotoId)">Delete this Photo</button>
		</div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p>Select a photo to upload</p>
			<input type="file" accept="image/*" @change="handleFile" ref="file">
			<button @click="uploadPhoto">Upload Photo</button>
		</div>
    <div>
      Your followers:
      <div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="user in followers" :key="user.ProfileId">
        <p>{{user.ProfileName}}</p>
      </div>
    </div>
    <div>
      You follow:
      <div class="col-md-3 justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 border-bottom" v-for="user in following" :key="user.ProfileId">
        <p>{{user.ProfileName}}</p>
      </div>
    </div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<button @click="goBack">Back to Homepage</button>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>

</style>
