<script setup>
import { RouterView } from 'vue-router'
</script>
<script>
export default {}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">WASAPhoto Login</a>
    <button class="navbar-toggler position-absolute collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#toggleButtons" aria-controls="toggleButtons" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse position-fixed top-0 end-0 mt-5" id="toggleButtons">
      <button class="btn btn-outline-dark d-block mb-2">Button 1</button>
      <button class="btn btn-outline-dark d-block">Button 2</button>
    </div>
  </header>

	<div class="container-fluid">
				<RouterView />
		</div>
</template>

<style>
</style>
