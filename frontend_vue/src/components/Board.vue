<template>
    <div class="board" id="board">
        <Octothorpe v-if="depth>0" :octDepth="depth - 1">
            <Octothorpe>
                <Octothorpe></Octothorpe>
            </Octothorpe>
        </Octothorpe>
<!--        <table class="oct-container">-->
<!--            <tr>-->
<!--                <td> <Octothorpe name="1" class="grid-cell bottom right"> <Cell /> </Octothorpe> </td>-->
<!--                <td> <Octothorpe name="2"> <Cell /> </Octothorpe>  </td>-->
<!--                <td> <Octothorpe name="3"> <Cell /> </Octothorpe> </td>-->
<!--            </tr>-->
<!--            <tr>-->
<!--                <td> <Octothorpe name="4"> <Cell /> </Octothorpe> </td>-->
<!--                <td> <Octothorpe name="5" > <Cell /> </Octothorpe> </td>-->
<!--                <td> <Octothorpe name="6"> <Cell /> </Octothorpe>  </td>-->
<!--            </tr>-->
<!--            <tr>-->
<!--                <td> <Octothorpe name="7" > <Cell /> </Octothorpe>  </td>-->
<!--                <td> <Octothorpe name="8" > <Cell /> </Octothorpe>  </td>-->
<!--                <td> <Octothorpe name="9" > <Cell /> </Octothorpe>  </td>-->
<!--            </tr>-->
<!--        </table>-->
    </div>
</template>

<script>
import Octothorpe from "./Octothorpe.vue";
// import Cell from "./Cell.vue";

// const depth = 3

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
