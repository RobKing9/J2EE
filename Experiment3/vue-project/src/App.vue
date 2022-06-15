<template>
  <div id="App">

<!--    <h1>HelloWord组件</h1>-->
    <router-view></router-view>
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld'

export default {
  name: 'App',
}
</script>

<style>
#App {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
