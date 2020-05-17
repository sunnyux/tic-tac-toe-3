<template>
  <button class="cell" :class="boardState" @click.once="markPlaced">
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
      coordID(){   // inside to outside
        // the next board will be defined from outside to inside
        let id = this.cellID
        let result = ""
        for(let i = this.$store.state.originalDepth + 1; i > 0; i--) {
          result = (id[0]).toString() + (id[1]).toString() + result
          id = id[2]
        }
        return result
      },
      boardID() {
        return this.coordID.substring(0, this.coordID.length - 2)
      },
      boardState() {
        console.log((this.$store.state.boardPlaying));
        if(this.$store.state.boardPlaying === "init" || this.$store.state.boardPlaying === this.boardID)
          if(this.$store.state.player === "X")
            return 'playerx';
          else //if(this.$store.state.player === "O")
            return 'playero';
        else
          return 'unplayable'
      }
    },
    methods: {
      markPlaced() {
        this.mark = this.$store.getters.getMark;
        this.$store.commit("markPlaced", this.coordID.substring(2))
      },
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
    pointer-events: none;
  }

  .playerx:focus, .playero:focus, .unplayable:focus {
    outline: 0;
    box-shadow: none!important;
  }

</style>
