<template>
<!--  <table :class="size" :id="id">-->
  <table :class="size">
    <tr v-for="row in 3" :key="row">
      <td v-for="column in 3" :key="row-column" :class="{borders: !moreNesting}">
        <div v-if="moreNesting" class="content">
          <Octothorpe :block-i-d="[row, column, blockID]"
                      :octDepth="octDepth-1" />
        </div>
        <div v-else>
          <Cell :cell-i-d="[row, column, blockID]" />
        </div>
      </td>
    </tr>
  </table>

</template>

<script>
  //TODO: if depth === 1 then special css
  //TODO: make the borders fucking conditional
  //TODO: need to figure out how to take html positions to the script as argument
  import Cell from "./Cell.vue";

  export default {
    name: "Octothorpe",
    props: {
      octDepth: {
        type: Number,
        required: true
      },
      blockID: {
        type: Array,
        required: true
      }
    },
    data() {
      return {
        lastTable: "lasttable",
      }
    },
    components: {
      Cell,
    },
    computed: {
      moreNesting() {
        return this.octDepth > 0;
      },
      id(){
        /*eslint-disable no-console*/
        console.log(("row: " + this.blockID.row + "  col: " + this.blockID.column));
        return this.blockID.row + "," + this.blockID.column
      },
      size() {
        if(!this.moreNesting)
          if(this.$store.state.originalDepth <= 1)
           return "size1";
          else
            return "size2"
        else
          return "size"
      },
      // color() {
      //   /*eslint-disable no-console*/
      //   // console.log((this.$store.state.boardPlaying));
      //   if(this.$store.state.boardPlaying === "init" || this.$store.state.boardPlaying === "X") {
      //     console.log("xplaying");
      //     return 'xplaying';
      //   }
      //   else if(this.$store.state.boardPlaying === "O")
      //     return 'oplaying';
      //   else
      //     return 'unplayable'
      // }
    },
    methods: {
      boardToPlay(board) {
        /* eslint-disable no-console */
        console.log(board)
        this.$store.commit("boardPlaying", board)
      }
    }

  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .content {
    position: absolute;
    margin: 0 0 0 0;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: gold;
    font-size: 90%;
    /*border-color: white;*/
  }
  td {
    width: 33.3333%;
    /*width: 5em;*/
    position: relative;
    padding: 0;
    margin: 0;
  }
  td:after {
    content: '';
    display: block;
    margin-top: 100%;
  }
  .borders {
    border: 2px solid white;
    margin: 0 0 0 0;
  }
  .size1 {
    width: 95%;
    height: 95%;
    left: 3%;
    top: 3%;
    margin-top: 3%;
    margin-left: 3%;
    table-layout: fixed;
    border-collapse: collapse;
  }
  .size2 {
    width: 98%;
    height: 98%;
    left: 1%;
    top: 1%;
    margin-top: 1%;
    margin-left: 1%;
    border-collapse: collapse;
  }
  .xplaying {
    background-color: lightsalmon;
  }
  .oplaying {
    background-color: cornflowerblue;
  }
  .unplayable {
    background-color: gold;
  }

</style>
