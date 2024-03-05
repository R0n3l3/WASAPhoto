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
				{
					photoId: 0,
					uploader: 0,
					uploadTime: "",
					likeNumber: 0,
					commentNumber: 0,
					image: "",
				}
			],
			profile: {
				profileId: 0,
				profileName: "",
				photoNumber: 0,
			}
		}
	},
	methods: {
		async getStream() {
			try {
				let response= await this.$axios.get("/users/"+this.username+"/profile", {
					params:
						{
							search: this.search
						},
					headers:
						{
							Authorization: "Bearer "+ this.token
						}
				})
				this.photoStream = response.data
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
		}
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
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
