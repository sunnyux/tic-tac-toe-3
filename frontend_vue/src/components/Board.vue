<template>
    <div class="board" id="board">
        <Octothorpe :oct-depth="octDepth"></Octothorpe>
    </div>
</template>

<script>
import Octothorpe from "./Octothorpe.vue";

export default {
    name: "Board",
    props: {
        octDepth: {
            type: Number,
            required: true,
        }
    },
    // data() {
    //     // boardPlaying: this.$store.boardPlaying;
    //     return {
    //         depth: 2
    //     };
    // },

    components: {
        Octothorpe,
        // Cell
    },

    methods: {
        doSomething() {
            alert("yes");
        }
    }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#board {
    /*border-radius: 10px;*/
    /*border: solid #533c00 5px;*/
    background-color: rgb(253, 255, 147);
    display: table;
    margin: 0 auto;
}

.board-playing {
    backgroung-color: rgb(60, 141, 191);
    margin: 0 auto;
}

.oct-container {
    margin: 3vw;
    border-collapse: collapse;
}

.bottom {
    border-bottom: solid black 5px;
}

.left {
    border-left: solid black 5px;
}

.right {
    border-right: solid black 5px;
}

</style>
