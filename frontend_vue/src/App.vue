<template>
  <div id="app">
    <Octothorpe :oct-depth="depth"></Octothorpe>
  </div>
</template>

<script>
// import Board from "./components/Board.vue";
import Octothorpe from "./components/Octothorpe";
export default {
  name: "app",
  components: {
    Octothorpe,
  },
  data() {
    return {
      depth: 0,
    }
  }
};
</script>
<style>
html {
  background-color: white;
}

#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: black;
}
</style>
