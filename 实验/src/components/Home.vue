<template>
  <div id="Home">
  	<h1>Dogs</h1>
  	<span><router-link to="/Home/one">First</router-link></span>
  	<span><router-link to="/Home/two">Second</router-link></span>
  	<span><router-link to="/Home/three">Third</router-link></span>
  	<br>
  	<router-view/>
  </div>
</template>

<script>
export default {
  name: 'Home'
}
</script>

<style scoped>
	.router-link-active{
	color:red;
	}
</style>
