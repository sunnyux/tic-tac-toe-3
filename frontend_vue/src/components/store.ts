import Vue from "vue";
import Vuex from "vuex";
Vue.use(Vuex);

export const store = new Vuex.Store({
  /* eslint-disable no-console */
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
    markPlaced(state, id) {
      if (state.player === "X") {
        state.player = "O";
      } else {
        state.player = "X"
      }
      // console.log((id));
      state.boardPlaying = id
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
