<template>
  <div id="Home">
<!--    <h1>我是首页</h1>-->
    <img src="../assets/logo.png">
    <router-link to="/index">跳转到首页</router-link>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: "Home"
}
</script>

<style scoped>

</style>
