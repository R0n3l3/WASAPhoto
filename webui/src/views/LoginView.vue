<script>
export default {
  components: {},
  data: function() {
    return {
      errormsg: null,
      username: "",
      profile: {
        profileId: 0,
        profileName: "",
      },
    }
  },
  methods: {
    async doLogin() {
      if (this.username==="") {
        this.errormsg="Username cannot be empty";
      } else {
        try {
          let response = await this.$axios.post("/session/", {username: this.username})
          this.profile=response.data
          localStorage.setItem("token", this.profile.profileId)
          localStorage.setItem("username", this.profile.profileName)
          this.$router.push({path:'/session'})
        } catch (e) {
          if (e.response && e.response.status===400) {
            this.errormsg="Form error, please check all fields and try again"
          } else if (e.response&&e.response.status===500) {
            this.errormsg="An internal error occurred"
          } else {
            this.errormsg=e.toString();
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
  <div class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
      <h1 class="h2">Welcome to WASAPhoto</h1>
    </div>
    <div class="input-group mb-3">
      <input type="text" id="username" v-model="username" class="form-control" placeholder="Insert a username" aria-label="Recipient's username" aria-describedby="basic-addon2">
      <div class="input-group-append">
        <button class="btn btn-success" type="button" @click="doLogin">Login</button>
      </div>
    </div>

    <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
</template>

<style>
</style>