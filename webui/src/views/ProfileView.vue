<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: "",
			detailedmsg: "",
			token: localStorage.getItem("token"),
			username: localStorage.getItem("username"),
			profileName: localStorage.getItem("searchName"),
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
			like: {
				likeId: 0,
				liker: 0,
				photoLiked: 0,
			},
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

		}
	},
	methods: {
		filteredPhotoData(photoId) {
			// Filter photoData based on photoId
			return this.photoData.filter(data => data.photoId === photoId);
		},

		async getData(photo) {
			let response = await this.$axios.get("/users/"+this.profileName+"/profile/photos/"+photo.PhotoId+"/comments/", {
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

			response = await this.$axios.get("/users/"+this.profileName+"/profile/photos/"+photo.PhotoId+"/likes/", {
				headers:
					{
						Authorization: "Bearer " + this.token
					}
			})
			if (response.data!==null) {
				this.likes=response.data
			}else {
				this.likes=[]
			}
			this.photoData.push({
				photoId: photo.PhotoId,
				photoComments: this.comments,
				photoLikes: this.likes,
			})
		},

		async refresh() {
			//Get my photos
			try {
				let response = await this.$axios.get("/users/" + this.profileName + "/profile/photos/", {
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
					this.getData(this.myPhotos[i])
				}
			}catch(e){
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
					this.errormsg = e.toString()
				}
			}
		},

		async checkLikes(photo) {
			this.photoData.forEach(p => {
				if (p.photoId === photo.PhotoId) {
					this.likes = p.photoLikes
				}
				if (this.likes === []) {
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
			try {
				await this.$axios.delete("/users/" + this.profileName + "/profile/photos/"+photo.PhotoId+"/likes/"+l.LikeId, {
					headers:
						{
							Authorization: "Bearer " + this.token
						}
				})
				photo.LikeNumber -= 1
				this.photoData = []
				this.refresh()
			} catch (e) {
				this.errormsg = e.toString()
			}
		},

		async commentPhoto(photo) {
			if (this.content==="") {
				this.errormsg="Please type a comment"
			}else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/users/"+this.profileName+"/profile/photos/"+photo.PhotoId+"/comments/", {
						content: this.content,
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
					photo.CommentNumber -= 1
					this.photoData = []
					this.refresh()
				} catch (e) {
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
			<h1 class="h2">{{this.profileName}} Profile</h1>
			You uploaded {{this.photoNumber}} photos
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
			<button @click="likePhoto(photo)">Add a like</button>
			<button @click="checkLikes(photo)">Remove a like</button>
			<input type="text" v-model="content" placeholder="Type a comment">
			<button @click="commentPhoto(photo)">Add a comment</button>
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
.resizable-image {
	max-width: 100%;
	max-height: 100px;
}
</style>
