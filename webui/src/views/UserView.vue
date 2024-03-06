<script>

import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function () {
		return {
			errormsg: null,
			detailedmsg: null,
			loading: false,
			username: localStorage.getItem("username"),
			newName: "",
			token: localStorage.getItem("token"),
			profile: {
				profileId: localStorage.getItem("searchId"),
				profileName: localStorage.getItem("searchName"),
				photoNumber: localStorage.getItem("searchPhoto"),
			},
			selectedFile: null
		}
	},
	methods: {
		async goBack() {
			localStorage.removeItem("searchId")
			localStorage.removeItem("searchName")
			localStorage.removeItem("searchPhoto")
			this.$router.push({path: "/session/"})
		},
		handleFileChange(event) {
			this.selectedFile=event.target.files[0]
		},

		async changeUsername() {
			if (this.newName==="") {
				this.errormsg="Please insert a name"
			} else {
				this.errormsg=null
				try {
					await this.$axios.put("/users/"+this.username, {
						new: this.newName
					}, {
						headers: {
							Authorization: "Bearer "+this.token
						}
					})
					this.username=this.newName
					this.profile.profileName=this.newName
					localStorage.setItem("username", this.newName)
					localStorage.setItem("searchName", this.newName)
					this.$router.push({path: "/users/" + this.newName + "/view/"})
				}
				catch(e) {
					this.errormsg = e.toString();
					this.detailedmsg = null;
				}
			}
				}

		},

		async upload() {
			if (!this.selectedFile) {
				this.errormsg="Please select a photo"
				return
			}
			const formData = new FormData();
			formData.append("link", this.selectedFile)

			try {
				let response = await this.$axios.post("/users/"+this.username+ "/profile/photos/", {
					formData
				}, {
					headers: {
						"Content-Type": "multipart/form-data",
						Authorization: "Bearer "+this.token
					}
				})
				this.selectedFile=null
			}catch(e) {
				console.log("error")
				this.errormsg=e.toString()
			}

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
			<input type="file" ref="fileInput" @change="handleFileChange">
			<button @click="upload">Upload Photo</button>
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
