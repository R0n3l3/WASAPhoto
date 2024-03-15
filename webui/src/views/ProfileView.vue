<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			username: localStorage.getItem("username"),
			loading: false,
			token: localStorage.getItem("token"),
			profile: {
				profileId: localStorage.getItem("searchId"),
				profileName: localStorage.getItem("searchName"),
				photoNumber: localStorage.getItem("searchPhoto"),
			},
			photos: [],
			comments: [],
			photosComments: [],
			commentContent: "",
			comment: {
				commentId: 0,
				commenter: 0,
				commentTime: "",
				content: "",
				photoComment: 0,
			},
			likes: [],
			like: {
				likeId: 0,
				photoLiked: 0,
				liker: 0,
			}
		}
	},
	methods: {
		async refresh() {
			try {
				let response = await this.$axios.get("/users/"+this.profile.profileName+"/profile/photos/", {
					headers: {
						Authorization: "Bearer " + this.token }
				})
				const photoArray=response.data
				if (photoArray===null) {
					this.detailedmsg="You haven't uploaded pictures"
				}else {
					this.detailedmsg=null
					this.photos=photoArray
					for (let i=0; i<this.photos.length; i++) {
						let response = await this.$axios.get("/users/"+this.profile.profileName+"/profile/photos/"+this.photos[i].PhotoId+"/comments/", {
							headers: {
								Authorization: "Bearer " + this.token }
						})
						this.comments=response.data
						response = await this.$axios.get("/users/"+this.profile.profileName+"/profile/photos/"+this.photos[i].PhotoId+"/likes/", {
							headers: {
								Authorization: "Bearer " + this.token }
						})
						this.likes=response.data
						this.photosComments.push({
							photoId: this.photos[i].PhotoId,
							comments: this.comments,
							likes: this.likes
						})
					}
					this.detailedmsg=null
				}
			}catch(e) {
				this.errormsg = e.toString();
				this.detailedmsg = null;
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
					await this.refresh()
				} catch (e) {
					this.errormsg = e.toString()
				}
			}
		},

		async likePhoto(item) {
			try {
				let response = await this.$axios.post("/users/"+item.Uploader+"/profile/photos/"+item.PhotoId+"/likes/", this.username,{
					headers: {
						Authorization: "Bearer " + this.token
					}
				})
				this.like=response.data
				item.LikeNumber=item.LikeNumber+=1
				this.photosComments.forEach(photo => {
					if (photo.photoId === item.PhotoId) {
						console.log(photo.likes)
						photo.likes.push(this.like)
					}
				})
				this.like=null
				await this.refresh()
			}catch(e){
				this.errormsg = e.toString()
			}
		},

		async goBack() {
			localStorage.removeItem("searchId")
			localStorage.removeItem("searchName")
			localStorage.removeItem("searchPhoto")
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
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">{{this.profile.profileName}} Profile</h1>
		</div>
		<div>
			<p>This user has uploaded {{this.profile.photoNumber}} photos.</p>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p> {{this.profile.profileName}} photos</p>
			<ul>
				<li v-for="(item, index) in photos" :key="index">
					<img :src="'data:image/*; base64,' + item.Image" alt="Uploaded Photo" class="resizable-image" />
					<p> Like Number: {{item.LikeNumber}}</p>
					<button @click="likePhoto(item)">Like Photo</button>
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
		<div>
			<button @click="goBack">Go to homepage</button>
		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
