<template>
    <!--    <div>-->
    <!--        <table class="oct-container">-->
    <!--            <tr>-->
    <!--                <td class="grid-cell top bottom left right">-->
    <!--                    <slot>Oct</slot>-->
    <!--                </td>-->
    <!--            </tr>-->
    <!--        </table>-->
    <!--    </div>-->
    <div>
        <slot>
            <table class="oct-container">
                <tr v-for="row in 3" :key="row">
                    <td v-for="column in 3" :key="row-column">
                        <div v-if="octDepth>0">
                            <Octothorpe :octDepth="octDepth-1"></Octothorpe>
                        </div>
                        <div v-else>
                            <Cell class="grid-cell top bottom left right"></Cell>
                        </div>
                    </td>
                </tr>
            </table>
        </slot>
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
    .oct-container {
        margin: 3vw;
        border-collapse: collapse;
    }

    .top {
        border-top: solid black 5px;
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

    .grid-cell {
        font-size: 7vw;
        position: relative;
        width: 10vw;
        height: 10vw
    }

    /*.grid-cell:after {*/
    /*    content: "";*/
    /*    display: block;*/
    /*    padding-bottom: 100%;*/
    /*}*/
</style>
