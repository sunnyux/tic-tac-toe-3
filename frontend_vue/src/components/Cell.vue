<template>
    <div :resize-text="{ratio: 0.6}">
      <button class="cell" ref="cell" :class="boardState" @click.once="markPlaced">
        <X icon-color="#f54c75" :style="{width: xWidth}" v-if="mark==='X'"></X>
        <O icon-color="#397fe6" :style="{width: oWidth}" v-else-if="mark==='O'"></O>
      </button>
    </div>
</template>

<script>
  /*eslint-disable no-console*/
  import X from "./icons/X";
  import O from "./icons/O";
  export default {
    name: "cell",
    components: {O, X},
    props: {
      cellID: {
        type: Array,
        required: true
      }
    },
    data() {
      return {
        mark: "",
        xWidth: 0,
        oWidth: 0,
      }
    },
    computed: {
      coordID() {   // inside to outside, xy-x'y'
        // the next board will be defined from outside to inside
        let id = this.cellID
        let result = ""
        for (let i = this.$store.state.originalDepth + 1; i > 0; i--) {
          result = (id[0]).toString() + (id[1]).toString() + "-" + result
          id = id[2]
        }
        return result
      },
      boardID() {
        return this.idFormatter(this.coordID, true)
      },
      boardState() {
        if (this.$store.state.boardPlaying === "init"
          || this.$store.state.boardPlaying === this.boardID)
          if (this.$store.state.player === "X")
            return 'playerx';
          else
            return 'playero';
        else
          return 'unplayable'
      }
    },
    methods: {
      getMarkSize() { // TODO: size is acceptable on my linux and mac rn for normal zoom (100%),
          // X and O goes inverse with initial zoom value, super small on phone
        this.xWidth = this.$refs.cell.clientWidth / 18 + 'vw';
        this.oWidth = this.$refs.cell.clientWidth / 17 + 'vw';
      },
      markPlaced() {
        let mark = this.$store.getters.getMark;
        if (mark === "X") this.mark = "X"
        else this.mark = "O"
        this.$store.commit("markPlaced",
                           this.idFormatter(this.coordID, false))
      },
      idFormatter(id, forBoard) { // if id is for boardID, then forBoard = true,
                                  // if it is for boardPlaying then false
        if (forBoard)
          return id.split('-').slice(0, this.$store.state.originalDepth).join("")
        else
          return id.split('-').slice(1).join("")
      }
    },
    mounted() {
      this.getMarkSize();
    }
  };
</script>

<style scoped>
  .cell {
    position: absolute;
    text-align: center;
    top: 0;
    bottom: 0;
    left: 0;
    right: 0;
    padding: 0 0 0 0;
    width: 100%;
    height: 100%;
    box-shadow: none;
    border: none;
  }

  .playerx {
    background: rgba(255, 187, 200, 0.53);
  }

  .playero {
    background-color: rgba(122, 191, 244, 0.42);
  }

  .unplayable {
    background-color: #ffdc52;
    pointer-events: none;
  }

  .playerx:focus, .playero:focus, .unplayable:focus {
    outline: 0;
    box-shadow: none !important;
  }

</style>
