import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

const store = new Vuex.Store({
  state: {
    letter: "X"
  },
  mutations: {
    ye(state) {
      state.letter = "y";
    }
  },
  getters: {
    getBoard(state) {
      return state.letter;
    }
  }
});

export default store;
