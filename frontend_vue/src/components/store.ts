import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

export const store = new Vuex.Store({
  state: {
    player: "X",
    originalDepth: 0,
    boardPlaying: "init",
    activeBoard: -1,
    winner: ""
  },
  mutations: {
    storeDepth(state, depth) {
      state.originalDepth = depth;
    },
    changeTurn(state) {
      if (state.player === "X") {
        state.player = "O";
      } else {
        state.player = "X"
      }
    },
    changeBoard(state, board) {
      state.boardPlaying = board;
    },

  },
  getters: {
    getMark(state) {
      return state.player;
    }
  }
});

export default store;
