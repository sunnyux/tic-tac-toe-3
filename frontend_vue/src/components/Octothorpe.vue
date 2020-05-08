<template>
        <table :class="{lasttable: !moreNesting}">
            <tr v-for="row in 3" :key="row">
                <td v-for="column in 3" :key="row-column" :class="{borders: !moreNesting}">
                    <div v-if="moreNesting" class="content">
                        <Octothorpe :block-i-d="{row: row, column: column, prevID: blockID}" :octDepth="octDepth-1"></Octothorpe>
                    </div>
                    <div v-else class="content">
                        <Cell :cell-i-d="{row: row, column: column, prevID: blockID}"></Cell>
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
                type: Object,
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
            nextBlockID(row, column) {
                /* eslint-disable no-console */
                console.log(this.blockID + "-" + row + "-" + column + "d" + this.octDepth-1);
                return this.blockID + "-" + row + "-" + column + "d" + this.octDepth-1;
                //d-1-1
            }
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
    }

    .last {
        /*border-collapse: collapse;*/
        position: absolute;
        left: 6%;
        /*right: 5%;*/
        top: 6%;
        /*bottom: 5%;*/
        margin: 0 0 0 0;
        padding: 0;
        width: 90%;
        height: 90%;
        background: white;
    }

    .lasttable {
        /*margin: 0 0 0 0;*/
        margin: auto;
        /*width: 8vw;*/
        /*padding-top: 3%;*/
        table-layout: fixed;
        left: 0;
        right: 0;
        top: 0;
        bottom: 0;
        width: 95%;
        height: 90%;
        border-collapse: collapse;
        border: 0;
    }


    /*.cell {*/
    /*    position: absolute;*/
    /*    !*margin: 0 0 0 0;*!*/
    /*    top: 0;*/
    /*    bottom: 0;*/
    /*    left: 0;*/
    /*    right: 0;*/
    /*    border-color: #2c3e50;*/
    /*}*/

    /*.center {*/
    /*    position: absolute;*/
    /*    left: 0;*/
    /*    top: 50%;*/
    /*    width: 100%;*/
    /*    !*text-align: center;*!*/
    /*    !*font-size: 18px;*!*/
    /*}*/

</style>
