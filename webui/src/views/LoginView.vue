<script>
import ErrorMsg from "../components/ErrorMsg.vue";

export default {
	components: {ErrorMsg},
	data: function() {
		return {
			errormsg: "",
      detailedmsg: "",
      username: "",
      id: 0,
		}
	},
	methods: {
		async doLogin () {
			if (this.username==="") {
				this.errormsg="Username cannot be empty"
			}
			else {
				this.errormsg=""
				try {
					let response = await this.$axios.post("/session/", this.username)
					this.id=response.data
					localStorage.setItem("token", this.id)
					localStorage.setItem("username", this.username)
          console.log(this.username)
					this.$router.push({path:"/session/"})
				}
				catch(e) {
						this.detailedmsg = e.toString();
						this.errormsg = "";
					}
			}
		}
	}
}
</script>

<template>
	<div>
		<div
			class="center-vertical d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="center-horizontal h2">Welcome to WASAPhoto!</h1>
		</div>
		<div class="justify-content-between align-items-center flex-wrap flex-md-nowrap pt-3 pb-2 mb-3">
      <p>Keep in touch with your friends by sharing photos of special moments, thanks to WASAPhoto! You can upload your
        photos directly from your PC, and they will be visible to everyone following you. </p>
      <div class="container">
        <img src="../components/images/free-images-1570092541.jpg" alt="photo" class="resizable-image with-border">
        <img src="../components/images/pexels-photo-236047-4170232318.jpeg" alt="photo" class="resizable-image with-border">
        <img src="../components/images/royalty-free-images-167608859.jpg" alt="photo" class="resizable-image with-border">
        <img src="../components/images/130-2687407817.jpg" alt="photo" class="resizable-image with-border">
      </div>
      </div>
    <div class="justify-content-between align-items-center flex-wrap flex-md-nowrap pt-3 pb-2 mb-3 border-bottom">
      <h6 class="center-horizontal center-vertical h9">Please insert a username to login or register:</h6>
      <div class="justify-content-between align-items-center flex-wrap flex-md-nowrap pt-3 pb-2 mb-3">
        <input type="text" v-model="username" placeholder="Username" class="center-horizontal form-control-lg">
        <button class="center-horizontal btn btn-outline-dark mb-3" @click="doLogin">Login</button>
      </div>
    </div>

		<ErrorMsg v-if="errormsg" :msg="errormsg" class="center-vertical center-horizontal"></ErrorMsg>
		<ErrorMsg v-if="detailedmsg" :msg="detailedmsg" class="center-vertical center-horizontal"></ErrorMsg>
	</div>
</template>

<style>

</style>
