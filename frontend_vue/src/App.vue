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
    // Board,
  },
  data() {
    return {
      depth: 2,
    }
  }
};
</script>
<style>
html {
  background-color: white;
}

/*#app {*/
/*  font-family: "Avenir", Helvetica, Arial, sans-serif;*/
/*  -webkit-font-smoothing: antialiased;*/
/*  -moz-osx-font-smoothing: grayscale;*/
/*  text-align: center;*/
/*  color: black;*/
/*}*/

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
