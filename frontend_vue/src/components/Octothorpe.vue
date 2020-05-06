<template>
    <table :class="{collapse: !nesting}">
        <tr v-for="row in 3" :key="row">
            <td v-for="column in 3" :key="row-column"  :class="{last: !nesting}">
                <div v-if="octDepth>0" class="content">
                    <Octothorpe :octDepth="octDepth-1"></Octothorpe>
                </div>
                <div v-else class="content">
                    <!--                        <div>X</div>-->
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
        components: {
            Cell,
        },
        computed: {
            nesting() {
                return this.octDepth > 0;
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
        /*border-color: white;*/
    }
    td {
        width: 33.33%;
        position: relative;
        padding: 0;
        margin: 0;
    }
    td:after {
        content: '';
        display: block;
        margin-top: 100%;
    }
    .collapse {
        border-collapse: collapse;
    }
    /*.last {*/
    /*    position: absolute;*/
    /*    padding: 0 0;*/
    /*    !*margin: 2% 2% 2% 2%;*!*/
    /*    top: 0;*/
    /*    bottom: 0;*/
    /*    left: 0;*/
    /*    right: 0;*/
    /*    !*background: white;*!*/
    /*    border-color: white;*/
    /*    !*border-collapse: collapse;*!*/
    /*}*/
    .last {
        border: 2px solid white;
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
