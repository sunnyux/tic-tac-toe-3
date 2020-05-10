<template>
  <button class="cell" :disabled="marked" @click="changeState">
    {{mark}}
  </button>
</template>

<script>
  // import store from "./store.ts";

  export default {
    name: "cell",

    props: {
      cellID: {
        type: Object,
        required: true
      }
    },

    data: function () {
      return {
        mark: ""
      }
    },
    computed: {
      marked() {
        return this.mark !== ""
      }
    },

    methods: {
      changeState() {
        this.mark = this.$store.getters.getMark;
        this.$store.commit("changeTurn")

        /* eslint-disable no-console */
        console.log(this.cellID)
      },
      nextBoard() {
        /* eslint-disable no-console */
        console.log(this.cellID)
      }
    }
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .cell {
    font-size: 100%;
    position: absolute;
    text-align: center;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    /*top: 50%;*/
    padding: 0 0 0 0;
    width: 95%;
    height: 95%;
    font-family: 'Gochi Hand', sans-serif;
    box-shadow: none;
    /*border: 2px solid white;*/
    border: none;
    background: gold;
  }

</style>
