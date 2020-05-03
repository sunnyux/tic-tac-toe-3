<template>
	<td class="cell">
		<div class="content" @click="changeState">
			{{mark}}
		</div>
	</td>
</template>

<script>
// import store from "./store.ts";

export default {
	name: "cell",

	props: ["name"],

	data: function () {
		return {
			mark: "X"
		}
	},

	// computed: {
	// 	getMark() {
	// 		return this.$store.getters.getMark
	// 	}
	// },

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
	font-size: 90%;
	position: relative;
	width: 50%;
	height: 10vw;
	width: 10vw;
	font-family: 'Gochi Hand', sans-serif;
}

.cell:after {
	content: "";
	display: block;
	padding-bottom: 100%;
}

.content {
	position: absolute;
	width: 100%;
	height: 100%;
}
</style>
