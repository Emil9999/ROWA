import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    farm_active: true,
    farm_info: {
      "1": {
        type: "lettuce",
        pos: {
          "1": {
            age: 3,
            harvestable: false
          },
          "2": {
            age: 10,
            harvestable: false
          },
          "3": {
            age: 17,
            harvestable: false
          },
          "4": {
            age: 24,
            harvestable: false
          },
          "5": {
            age: 31,
            harvestable: false
          },
          "6": {
            age: 38,
            harvestable: false
          }
        }
      },
      "2": {
        type: "lettuce",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "3": {
        type: "basil",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "4": {
        type: "basil",
        pos: {
          1: false,
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: false
          }
        }
      },
      "5": {
        type: "lettuce",
        pos: {
          1: {
            age: 3,
            harvestable: false
          },
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: true
          }
        }
      },
      "6": {
        type: "lettuce",
        pos: {
          1: false,
          2: {
            age: 10,
            harvestable: false
          },
          3: {
            age: 17,
            harvestable: false
          },
          4: {
            age: 24,
            harvestable: false
          },
          5: {
            age: 31,
            harvestable: false
          },
          6: {
            age: 42,
            harvestable: false
          }
        }
      },
    }
  },
  mutations: {
    CHANGE_DASH_STATE(state){
      state.farm_active = !state.farm_active
    }
  },
  actions: {
    change_dash_state({commit}){
      commit("CHANGE_DASH_STATE")
    }
  },
  modules: {
  }
})
