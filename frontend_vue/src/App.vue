<template>
  <div id="app">
    <Board :oct-depth="depth"></Board>
  </div>
</template>

<script>
import Board from "./components/Board.vue";
// import Octothorpe from "./components/Octothorpe";
export default {
  name: "app",
  components: {
    // Octothorpe,
    Board,
  },
  data() {
    return {
      depth: 1,
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
