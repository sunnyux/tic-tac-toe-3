<template>
  <button class="cell" :class="color" @click.once="markPlaced">
    {{mark}}
  </button>
</template>

<script>
  //TODO: Disable click effect
  // import store from "./store.ts";
  /*eslint-disable no-console*/

  export default {
    name: "cell",
    props: {
      cellID: {
        type: Array,
        required: true
      }
    },
    data() {
      return {
        mark: ""
      }
    },
    computed: {
      marked() {
        return this.mark !== ""
      },
      uniqueID(){   // inside to outside
        // the next board will be defined from outside to inside
        return "c" + this.cellID.flat(this.$store.state.originalDepth).join("");
      },
      color() {
        console.log((this.$store.state.boardPlaying));
        if(this.$store.state.boardPlaying === "init" || this.$store.state.boardPlaying === "X")
          return 'playerx';
        else if(this.$store.state.boardPlaying === "O")
          return 'playero';
        else
          return 'unplayable'
      }
    },
    methods: {
      markPlaced() {
        this.mark = this.$store.getters.getMark;
        this.$store.commit("markPlaced", this.cellID)

        //TODO (rule + display): change board color after play
        /* eslint-disable no-console */
        // console.log(this.cellID)
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
    width: 100%;
    height: 100%;
    font-family: 'Gochi Hand', sans-serif;
    box-shadow: none;
    /*border: 2px solid white;*/
    border: none;
  }

  .playerx {
    background: lightsalmon;
  }

  .playero {
    background-color: cornflowerblue;
  }

  .unplayable {
    background-color: gold;
  }

</style>
