import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    player: "X",
    boardPlaying: 0,
    winner: ""
  },
  mutations: {
    changeTurn(state) {
      if (state.player === "X") {
        state.player = "O";
      } else {
        state.player = "X"
      }
    },
    changeBoard(state, board) {
      state.boardPlaying = board;
    }
  },
  getters: {
    getMark(state) {
      return state.player;
    }
  }
});

export default store;
