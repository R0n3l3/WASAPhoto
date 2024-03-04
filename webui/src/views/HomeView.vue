<script>

export default {
	components: {},
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
				let response = await this.$axios.get("/users/" + this.username + "/profile/", {
					headers: {
						Authorization: "Bearer " + this.token
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
					this.profile = response.data


				} catch (e) {
					if (e.response && e.response.status === 404) {
						this.detailedmsg = "Your stream is empty"
					} else {
						this.detailedmsg = e.toString()
					}

				}
			}
		},
		mounted() {
			this.getStream()
		}
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
