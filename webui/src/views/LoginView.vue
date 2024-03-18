<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
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
				this.errormsg="Username cannot be empty"
			}
			else {
				this.errormsg=null
				try {
					let response = await this.$axios.post("/session/?username="+this.username)
					this.id=response.data
					localStorage.setItem("token", this.id)
					localStorage.setItem("username", this.username)
					this.$router.push({path:"/session/"})
				}
				catch(e) {
						this.errormsg = e.toString();
						this.detailedmsg = null;
					}
			}
		}
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
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
