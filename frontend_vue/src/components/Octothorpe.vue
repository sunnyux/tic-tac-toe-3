<template>
  <div>
    <table :class="size">
      <tr v-for="row in boardSize" :key="row">
        <td v-for="column in boardSize" :key="row+column" :class="{borders: !moreNesting}">
          <div v-if="moreNesting" class="content">
            <Octothorpe :block-i-d="[row, column, blockID]"
                        :octDepth="octDepth-1" :board-size="parseInt(boardSize)"/>
          </div>
          <div v-else>
            <Cell :cell-i-d="[row, column, blockID]"/>
          </div>
        </td>
      </tr>
    </table>
    <div :resize-text="{ratio: 1}">
      <div v-if="completed" class="winner">{{winner}}</div>
    </div>
  </div>
</template>

<script>
  import Cell from "./Cell.vue";

  export default {
    name: "Octothorpe",
    props: {
      octDepth: {
        type: Number,
        required: true
      },
      boardSize: {
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
        winner: "×",
      }
    },
    components: {
      Cell,
    },
    computed: {
      completed() {
        return false
        // return this.octDepth === 0;
      },
      moreNesting() {
        return this.octDepth > 0;
      },
      size() {
        if (!this.moreNesting)
          if (this.$store.state.originalDepth <= 1)
            return "size1";
          else
            return "size2"
        else
          return "size"
      },
    },
  };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  .winner {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 100%;
    height: 100%;
    font-size: 30vw;
    color: #eddbbf;
    background-color: #846c00;
    opacity: 50%;
    display: flex;
    align-items: center;
    justify-content: center;
    transform: translate(-50%, -50%);
    -ms-transform: translate(-50%, -50%);
  }

  table {
    border-style: hidden;
  }

  .content {
    position: absolute;
    margin: 0 0 0 0;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    background: #ffee9c;
    /*font-size: 90%;*/
  }

  td {
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
    left: 2.5%;
    top: 2.5%;
    margin-top: 2.5%;
    margin-left: 2.5%;
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

</style>
