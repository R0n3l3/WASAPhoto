<script setup>
import { RouterView } from 'vue-router'
import {useRouter} from "vue-router";
import {ref, watchEffect} from "vue";

const router = useRouter()
const showNavbar = ref(true)
const username = ref(localStorage.getItem("username"))
const token = ref(localStorage.getItem("token"))

watchEffect(() => {
    const route = router.currentRoute.value
    showNavbar.value = route.path !== '/';
    username.value=localStorage.getItem("username");
    token.value=localStorage.getItem("token")

  })
</script>
<script>
import ErrorMsg from "./components/ErrorMsg.vue";

export default {
  components: {ErrorMsg},
  data: function () {
    return {
      errormsg: "",
      detailedmsg: "",
      profile: {
        profileId: 0,
        profileName: "",
        photoNumber: 0,
      },
    }
  },
  methods: {
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
        localStorage.setItem("searchId", this.profile.profileId)
        localStorage.setItem("searchPhoto", this.profile.photoNumber)
        console.log(this.profile.profileName)
        console.log(localStorage.getItem("searchName"))
        this.$router.push({path: "/users/" + this.profile.profileName + "/view/"})
      } catch (e) {
        if (e.response && e.response.status === 404) {
          this.errormsg = "User not found"
        } else {
          this.errormsg = e.toString()
        }
      }
    },

    async homepage() {
      localStorage.setItem("searchName", "")
      localStorage.setItem("searchId", "")
      localStorage.setItem("searchPhoto", "")
      this.$router.push({path: "/session/"})
    },
  },
}


</script>

<template>

	<header v-if="showNavbar" class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto Login</a>
    <button class="navbar-toggler position-absolute collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#toggleButtons" aria-controls="toggleButtons" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse position-fixed top-0 end-0 mt-5" id="toggleButtons">
      <button @click="getMyProfile" class="btn btn-outline-dark d-block mb-2">Your Profile</button>
      <button @click="homepage" class="btn btn-outline-dark d-block">Homepage</button>
    </div>
  </header>

	<div class="container-fluid">
				<RouterView />
		</div>
</template>

<style>
</style>
