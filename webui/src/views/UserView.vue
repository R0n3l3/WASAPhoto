<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			commentContent: "",
			errormsg: null,
			detailedmsg: null,
			loading: false,
			username: localStorage.getItem("username"),
			newName: "",
			token: localStorage.getItem("token"),
			profile: {
				profileId: localStorage.getItem("searchId"),
				profileName: localStorage.getItem("searchName"),
				photoNumber: parseInt(localStorage.getItem("searchPhoto")),
			},
			myphotos: [
			],
			myphotoscomments:[
			],
			photo : {
				photoId: 0,
				uploader: "",
				image: "",
				uploadTime: "",
				likeNumber: 0,
				commentNumber: 0,
				comments: [],
			},
			image: null,
			comments: [],
			comment: {
				commentId: 0,
				commenter: 0,
				commentTime: "",
				content: "",
				photoComment: 0,
			}
		}
	},
	methods: {
		async refresh() {
			try {
				let response = await this.$axios.get("/users/"+this.username+"/profile/photos/", {
				headers: {
					Authorization: "Bearer " + this.token }
				})
				const photoArray=response.data
				if (photoArray===null) {
					this.detailedmsg="You haven't uploaded pictures"
				}else {
					this.detailedmsg=null
					this.myphotos=photoArray
					for (let i=0; i<this.myphotos.length; i++) {
						let response = await this.$axios.get("/users/"+this.username+"/profile/photos/"+this.myphotos[i].PhotoId+"/comments/", {
							headers: {
								Authorization: "Bearer " + this.token }
						})
						this.comments=response.data
						this.myphotoscomments.push({
						photoId: this.myphotos[i].PhotoId,
						comments: this.comments
					})
					}
					this.detailedmsg=null
				}
		}catch(e) {
				this.errormsg = e.toString();
				this.detailedmsg = null;
			}},

		async goBack() {
			localStorage.removeItem("searchId")
			localStorage.removeItem("searchName")
			localStorage.removeItem("searchPhoto")
			this.$router.push({path: "/session/"})
		},

		async changeUsername() {
			if (this.newName === "") {
				this.errormsg = "Please insert a name"
			} else {
				this.errormsg = null
				try {
					await this.$axios.put("/users/" + this.username, {
						new: this.newName
					}, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.username = this.newName
					this.profile.profileName = this.newName
					localStorage.setItem("username", this.newName)
					localStorage.setItem("searchName", this.newName)
					this.refresh()
				} catch (e) {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
		},

		async handleFileChange() {
			this.image=this.$refs.file.files[0]
		},

		async upload() {
			if (this.image===null) {
				this.errormsg="Please select a photo"
			}else {
				this.errormsg=null
				try {
					let response = await this.$axios.post("/users/" + this.username + "/profile/photos/", this.image, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.photo=response.data
					this.myphotos.unshift(this.photo)
					this.myphotoscomments.push({
						photoId: this.photo.photoId,
						comments: []
					})
					localStorage.setItem("searchPhoto", this.profile.photoNumber+1)
					this.profile.photoNumber=this.profile.photoNumber+1
					this.photo=null
					this.$refs.file.value = ''
					this.refresh()
				} catch (e) {
					this.errormsg = e.toString()
				}
			}
		},

		async deletePhoto(id, index) {
			try {
				await this.$axios.delete("/users/"+this.username+"/profile/photos/"+id, {
					headers: {
						Authorization: "Bearer " + this.token
					}
				})
				this.myphotos.splice(index, 1)
				localStorage.setItem("searchPhoto", this.profile.photoNumber-1)
				this.profile.photoNumber=this.profile.photoNumber-1
				this.myphotoscomments=null
				this.refresh()
			}catch(e){
				this.errormsg = e.toString()
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
					this.myphotoscomments.forEach(photo => {
						if (photo.photoId === item.PhotoId) {
							photo.comments.push(this.comment);
						}
					});
					this.commentContent=null
					await this.refresh()
				} catch (e) {
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
					this.myphotoscomments.forEach(p => {
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
			<h1 class="h2">{{this.profile.profileName}} Profile</h1>
		</div>
		<div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p>This user has uploaded {{this.profile.photoNumber}} photos.</p>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p>Please select a photo to upload</p>
			<input type="file" accept="image/*" @change="handleFileChange" ref="file">
			<button @click="upload">Upload Photo</button>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p>Change your username</p>
			<input type="text" v-model="newName">
			<button @click="changeUsername">Update Name</button>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p> Your photos</p>
			<ul>
				<li v-for="(item, index) in myphotos" :key="index">
					<img :src="'data:image/*; base64,' + item.Image" alt="Uploaded Photo" class="resizable-image" />
					<p> Like Number: {{item.LikeNumber}}</p>
					<p> Comment Number: {{item.CommentNumber}}</p>
					<ul>
						<li v-if="myphotoscomments[index]" v-for="(comment, i) in myphotoscomments[index].comments" :key="comment.CommentId">
							<p>{{comment.Commenter}}: {{comment.Content}}</p>
							<button @click="deleteComment(comment, i, item)">Delete comment</button>
						</li>
					</ul>
					<input type="text" v-model="commentContent">
					<button @click="commentPhoto(item)">Add a comment</button>
					<p> Date of Upload: {{item.UploadTime}}</p>
					<button @click="deletePhoto(item.PhotoId, index)"> Remove Photo</button>
				</li>
			</ul>
		</div>
		<button @click="goBack">Go to homepage</button>
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
