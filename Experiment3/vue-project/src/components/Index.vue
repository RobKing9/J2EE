<template>
  <div>
    <img src="../assets/logo.png">
    <h1 style="position: absolute;left: 730px; top: 215px;">首页</h1>
    <!--  <router-view></router-view>-->
    <ul style="position: absolute;left: 655px;top: 310px">
      <span><router-link to="/index/login">登录页面</router-link></span>
      <span><router-link to="/index/register">注册页面</router-link></span>
    </ul>
    <router-view></router-view>
  </div>

</template>

<script>
export default {
  name: "Index"
}
</script>

<style scoped>
  .router-link-active{
    color: red;
  }
</style>
