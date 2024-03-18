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
		<div class="justify-content-between align-items-center flex-wrap flex-md-nowrap pt-3 pb-2 mb-3 border-bottom">
      <p>Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can upload your
      photos directly from your PC, and they will be visible to everyone following you. </p>
      <img src="../components/images/free-images-1570092541.jpg" alt="photo" class="resizable-image with-border">
      <img src="../components/images/pexels-photo-236047-4170232318.jpeg" alt="photo" class="resizable-image with-border">
      <img src="../components/images/royalty-free-images-167608859.jpg" alt="photo" class="resizable-image with-border">
    </div>
      <h6 class="h9">Please insert a username to login or register:</h6>
		<input type="text" v-model="username"
			placeholder="Username">
			<button @click="doLogin">Login</button>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg"></ErrorMsg>
	</div>
</template>

<style>

</style>
