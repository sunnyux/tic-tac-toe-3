<template>
  <div class="xo" @click="changeState">{{this.$store.getters.getBoard}}</div>
</template>

<script>
import store from "./store.ts";
export default {
  name: "XO",

  methods: {
    changeState() {
      this.$store.commit("ye");
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
</style>
