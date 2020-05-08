<template>
<!--	<td>-->
		<button class="cell" @click="changeState">
			{{mark}}
		</button>
<!--	</td>-->
</template>

<script>
// import store from "./store.ts";

export default {
	name: "cell",

	props: {
		cellID: {
			type: Object,
			requied: true
		}
	},

	data: function () {
		return {
			mark: ""
		}
	},

	methods: {
		changeState() {
			/* eslint-disable no-console */
			console.log("afs")
			// this.$store.commit("ye");
			this.mark = this.$store.getters.getMark;
			this.$store.commit("changeTurn")
		}
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
	width: 95%;
	height: 95%;
	font-family: 'Gochi Hand', sans-serif;
	box-shadow: none;
	/*border: 2px solid white;*/
	border: none;
	background: gold;
}

/*.center {*/
/*	position: absolute;*/
/*	left: 0;*/
/*	top: 50%;*/
/*	width: 100%;*/
/*	text-align: center;*/
/*	font-size: 18px;*/
/*}*/

/*.cell:after {*/
/*	content: "";*/
/*	display: block;*/
/*	padding-bottom: 100%;*/
/*}*/

/*table {*/
/*	width: 90%;*/
/*}*/
/*td {*/
/*	background: gold;*/
/*}*/
/*td:after {*/
/*	display: block;*/
/*	margin-top: 90%;*/
/*}*/

/*.content {*/
/*	position: absolute;*/
/*	width: 100%;*/
/*	height: 100%;*/
/*}*/

/*.content {*/
/*	position:absolute;*/
/*	height:80%; !* = 100% - 2*10% padding *!*/
/*	width:90%; !* = 100% - 2*5% padding *!*/
/*	padding: 10% 5%;*/
/*}*/
</style>
