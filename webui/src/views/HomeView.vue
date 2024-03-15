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
		}
	},
	methods: {
		async getStream() {
			try {
				let response= await this.$axios.get("/users/"+this.username+"/profile", {
					headers:
						{
							Authorization: "Bearer "+ this.token
						}
				})
				this.photoStream = response.data
				this.detailedmsg=null
				for (let i=0; i<this.photoStream; i++) {
					let response = await this.$axios.get("/users/"+this.username+"/profile/photos/"+this.photoStream[i].PhotoId+"/comments/", {
						headers: {
							Authorization: "Bearer " + this.token }
					})
					this.comments=response.data
					this.photosComments.push({
						photoId: this.photoStream[i].PhotoId,
						comments: this.comments
					})
				}
				this.detailedmsg=null
			} catch (e) {
				if (e.response && e.response.status === 404) {
					this.detailedmsg = "Your stream is empty"
				} else {
					this.detailedmsg = e.toString()
				}
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
					if (decodedprofile.ProfileName===this.username) {
						this.errormsg="You cannot search yourself! Please use the 'My profile' button"
					} else {
						this.profile.profileName = decodedprofile.ProfileName
						this.profile.profileId = decodedprofile.ProfileId
						this.profile.photoNumber = decodedprofile.PhotoNumber
						localStorage.setItem("searchName", this.profile.profileName)
						localStorage.setItem("searchId", this.profile.profileName)
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
				localStorage.setItem("searchId", this.profile.profileName)
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

		async deleteComment(comment, index, photo) {
			if (this.username!==comment.Commenter) {
				this.errormsg="You cannot delete this comment"
			}else {
				try {
					await this.$axios.delete("/users/" + this.username + "/profile/photos/" + photo.PhotoId + "/comments/" + comment.CommentId, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					photo.CommentNumber = photo.CommentNumber -= 1
					this.photosComments.forEach(p => {
						if (p.photoId === photo.PhotoId) {
							p.comments.splice(index, 1)
						}
					})
					this.refresh()
				} catch (e) {
					this.errormsg = e.toString()
				}
			}
		},

		async commentPhoto(item) {
			if (this.commentContent==="") {
				this.errormsg = "Please type a comment"
			}else {
				try {
					let response = await this.$axios.post("/users/" + this.username + "/profile/photos/" + item.PhotoId + "/comments/", this.commentContent, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.comment=response.data
					item.CommentNumber = item.CommentNumber+=1
					this.photosComments.forEach(photo => {
						if (photo.photoId === item.PhotoId) {
							photo.comments.push(this.comment);
						}
					});
					this.commentContent=null
					await this.getStream()
				} catch (e) {
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
			<ul>
				<li v-for="(item, index) in photoStream" :key="index">
					<img :src="'data:image/*; base64,' + item.Image" alt="Uploaded Photo" class="resizable-image" />
					<p> Like Number: {{item.LikeNumber}}</p>
					<p> Comment Number: {{item.CommentNumber}}</p>
					<ul>
						<li v-if="photosComments[index]" v-for="(comment, i) in photosComments[index].comments" :key="comment.CommentId">
							<p>{{comment.Commenter}}: {{comment.Content}}</p>
							<button @click="deleteComment(comment, i, item)">Delete comment</button>
						</li>
					</ul>
					<input type="text" v-model="commentContent">
					<button @click="commentPhoto(item)">Add a comment</button>
					<p> Date of Upload: {{item.UploadTime}}</p>
				</li>
			</ul>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
