<script>
export default {
	data: function() {
		return {
			errormsg: null,
			username: "",
			id: 0,
			detailedmsg: null,
		}
	},
	methods: {
		async doLogin () {
			if (this.username==="") {
				this.errormsg="You must insert a username!"
			} else {
				try {
					let response = await this.$axios.post("/session/", {username: this.username})
					this.id=response.data
					localStorage.setItem("token", this.id)
					localStorage.setItem("username", this.username)
					this.$router.push({path:'/session'})
				} catch (e) {
					if (e.response && e.response.status === 400) {
						this.errormsg = "Form error, please check all fields and try again. If you think that this is an error, write an e-mail to us.";
						this.detailedmsg = null;
					} else if (e.response && e.response.status === 500) {
						this.errormsg = "An internal error occurred. We will be notified. Please try again later.";
						this.detailedmsg = e.toString();
					} else {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
				}
			}
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
			<h1 class="h2">Welcome to WASAPhoto!</h1>
		</div>
		<div>
			<p>Please insert your username to login or register:</p>
		<input type="text" v-model="username"
			placeholder="Username">
			<button @click="doLogin">Login</button>
		</div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
