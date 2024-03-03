<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function() {
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
		async getStream () {
			try {
					let response = await this.$axios.get("/users/" + this.username + "/profile/following/", {
						headers: {
							Authorization: "Bearer " + this.token
			}
			})
					this.photoStream=response.data

				} catch (e) {
				if (e.response && e.response.status === 404) {
					this.detailedmsg = "Your stream is empty";
				}
				if (e.response && e.response.status === 405) {
					this.detailedmsg = "Error";


				}
			}
			},
		async getUser() {
			if (this.search==="") {
				this.errormsg="You must insert a username!"
			}
			else {
				try {
					let response = await this.$axios.get("/users/" + this.username + "/profile/?username=" + this.search, {
						headers: {
							Authorization: "Bearer " + this.token
						}
					})
					this.profile = response.data
					localStorage.setItem("profile", this.profile)
					this.$router.push ({path:"/users/"+this.profile.profileId+"/profile/"})
				} catch (e) {
				if (e.response && e.response.status===404) {
					this.errormsg= "Profile not found"
				}}
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

		</div>
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
