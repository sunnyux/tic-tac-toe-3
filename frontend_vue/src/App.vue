<template>
  <div id="app">
    <Octothorpe :block-i-d="initialID" :oct-depth="depth" :board-size="size"></Octothorpe>
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
        depth: 2,
        size: 3, // the number of rows and columns of the board
      }
    },
    created() {
      this.$store.commit("storeDepth", this.depth)
      this.$store.commit("storeSize", this.size)
    },
    computed: {
      initialID() {
        return ["d"]
      }
    }
  };
</script>
<style>
  html {
    background-color: white;
  }

  #app {
    font-family: Avenir, Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
    margin-top: 30px;
    font-size: 2vw;
  }

  table {
    margin: auto;
    /*width: 8vw;*/
    table-layout: fixed;
    min-width: 95%;
    padding: 3% 0.1%;
    /*border-color: black;*/
    border-collapse: separate;
  }
</style>

