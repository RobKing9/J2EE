<template>
  <div id="login">
    <h3>{{msg}}</h3>
    <br>
    <img src="../images/login.png">
  </div>
</template>

<script>
export default {
  name: "Login",
  data(){
    return {
      msg:'密码登录'
    }
  }
}
</script>

<style scoped>

</style>
