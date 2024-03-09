<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			loading: false,
			uploaded: null,
			username: localStorage.getItem("username"),
			newName: "",
			token: localStorage.getItem("token"),
			profile: {
				profileId: localStorage.getItem("searchId"),
				profileName: localStorage.getItem("searchName"),
				photoNumber: localStorage.getItem("searchPhoto"),
			},
			photo : {
				photoId: 0,
				uploader: "",
				image: "",
				uploadTime: "",
				likeNumber: 0,
				commentNumber: 0,
			},
			image: null
		}
	},
	methods: {
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
					this.$router.push({path: "/users/" + this.newName + "/view/"})
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
					this.uploaded='data:image/*; base64, '+this.photo.Image

					this.image = null

				} catch (e) {
					console.log("error")
					this.errormsg = e.toString()
				}
			}
		},

		},


	mounted() {
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
			<img v-if="uploaded" :src="uploaded" alt="Example"/>
		</div>
		<div class="d-flex flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<p>Change your username</p>
			<input type="text" v-model="newName">
			<button @click="changeUsername">Update Name</button>
		</div>
		<button @click="goBack">Go to homepage</button>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
