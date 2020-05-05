<template>
        <table class="square">
            <tr v-for="row in 3" :key="row">
                <td v-for="column in 3" :key="row-column">
                    <div v-if="octDepth>0">
                        <Octothorpe :octDepth="octDepth-1"></Octothorpe>
                    </div>
                    <div v-else>
                        <Cell></Cell>
                    </div>
                </td>
            </tr>
        </table>
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
        },
        // },
        // data() {
        //     return {
        //         octDepth: 0,
        //         isTop: 'top',
        //         isBottom: 'bottom',
        //         isLeft: 'left',
        //         isRight: 'right',
        //     }
        // },
        components: {
            Cell,
        },
        computed: {
            getClasses(row, column) {
                /* eslint-disable no-console */
                let result = []
                if (row === 0) result.push(this.isTop)
                if (row === 2) result.push(this.isBottom)
                if (column === 0) result.push(this.isLeft)
                if (column === 2) result.push(this.isRight)
                if (row === 1 && column === 1) result.push(this.isTop, this.isBottom, this.isLeft, this.isRight)
                console.log(row);
                return result
            },
        },
        methods: {
            log() {
                console.log("called");
            },
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
    /*.oct-container {*/
    /*    margin: 3vw;*/
    /*    border-collapse: collapse;*/
    /*}*/

    .top {
        border-top: solid black 1px;
    }

    .bottom {
        border-bottom: solid black 1px;
    }

    .left {
        border-left: solid black 1px;
    }

    .right {
        border-right: solid black 1px;
    }

    .grid-cell {
        font-size: 7vw;
        position: relative;
        width: 10vw;
        height: 10vw
    }

    table {
        border: 1px solid black;
        width: 10vw;
        height: 30%;
        border-collapse: collapse;
    }


    /*.grid-cell:after {*/
    /*    content: "";*/
    /*    display: block;*/
    /*    padding-bottom: 100%;*/
    /*}*/

    div {
        width: 30%;
        padding-bottom: 30%; /* = width for a square aspect ratio */
    }

    .square {
        float:left;
        position: relative;
        width: 30%;
        padding-bottom: 30%; /* = width for a 1:1 aspect ratio */
        margin:1.66%;
        overflow:hidden;
    }

    .content {
        position:absolute;
        height:80%; /* = 100% - 2*10% padding */
        width:90%; /* = 100% - 2*5% padding */
        padding: 10% 5%;
    }

</style>
