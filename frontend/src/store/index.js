import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    farm_active: true,
    batch_active: true,
    admin_active: true,
    yPos_plantInfo: 0,
    yPos_statInfo: 0,
    to_farm:  {
      module: 0,
      position: 0,
      plant_type: "Lettuce"
    },
    farm_info: {
      "1": {
        type: "lettuce",
        pos: [ {
            age: 3,
            harvestable: true
          },{
            age: 10,
            harvestable: true
          },{
            age: 17,
            harvestable: true
          },{
            age: 24,
            harvestable: true
          }, {
            age: 31,
            harvestable: false
          },{
            age: 38,
            harvestable: false
          }
        ],
      },
      "2": {
        type: "lettuce",
        pos: [ {
            age: 3,
            harvestable: false
          },
          {
            age: 10,
            harvestable: false
          },
          {
            age: 17,
            harvestable: false
          },
          {
            age: 24,
            harvestable: false
          },
         {
            age: 31,
            harvestable: false
          },
           {
            age: 42,
            harvestable: true
          }
        ]
      },
      "3": {
        type: "basil",
        pos:[
          {
            age: 3,
            harvestable: false
          },{
            age: 10,
            harvestable: false
          },{
            age: 17,
            harvestable: false
          },{
            age: 24,
            harvestable: false
          },{
            age: 31,
            harvestable: false
          },{
            age: 42,
            harvestable: true
          }
        ]
      },
      "4": {
        type: "basil",
        pos: [
          {
            age: null,
            harvestable: null
          }, {
            age: 10,
            harvestable: false
          },{
            age: 17,
            harvestable: false
          },{
            age: 24,
            harvestable: false
          }, {
            age: 31,
            harvestable: false
          }, {
            age: 42,
            harvestable: false
          }
        ]
      },
      "5": {
        type: "lettuce",
        pos: [
           {
            age: 3,
            harvestable: false
          }, {
            age: 10,
            harvestable: false
          }, {
            age: 17,
            harvestable: false
          },{
            age: 24,
            harvestable: false
          }, {
            age: 31,
            harvestable: false
          },{
            age: 42,
            harvestable: true
          }
        ]
      },
      "6": {
        type: "lettuce",
        pos: [{
            age: null,
            harvestable: null
          }, {
            age: 10,
            harvestable: false
          }, {
            age: 17,
            harvestable: false
          },{
            age: 24,
            harvestable: false
          }, {
            age: 31,
            harvestable: false
          }, {
            age: 42,
            harvestable: false
          }
        ]
      },
    }
  },
  mutations: {
    InsertFarming(state, load){
      state.to_farm.module = load.m
      state.to_farm.position = load.p
      state.to_farm.plant_type = load.t
    },
    CHANGE_YPOS_PLANTINFO(state, load){
      state.yPos_plantInfo = load
    },
    CHANGE_YPOS_STATINFO(state, load){
      state.yPos_statInfo = load
    },
    CLEAR_FARMING(state){
      state.to_farm.module = 0
      state.to_farm.position = 0
      state.to_farm.plant_type = "null"
    },
    CHANGE_DASH_STATE(state){
      state.farm_active = !state.farm_active
    },
    CHANGE_BATCH_STATE(state){
      state.batch_active = !state.batch_active
    },
    CHANGE_ADMIN_STATE(state){
      state.admin_active = !state.admin_active
    },
    FarmUpdate(state,  data){
      console.log(data)
      state.farm_info[data.moduleNumber].type = data[0].plant_type.toLowerCase()
      for(var i = 0;i<6;i++){
        state.farm_info[data.moduleNumber].pos[i].harvestable = null
        state.farm_info[data.moduleNumber].pos[i].age = 0
      }
      for(var j = 0;j<6;j++){
        if(data[j]!== undefined){
          state.farm_info[data.moduleNumber].pos[data[j].plant_position-1].harvestable = data[j].harvestable
          state.farm_info[data.moduleNumber].pos[data[j].plant_position-1].age = data[j].age
          console.log(state.farm_info[data.moduleNumber].pos[j+1].harvestable)
        }
      }
      

    }
  },
  actions: {
    insertFarming: ({commit}, payload) => {
      commit('InsertFarming',  payload )
    },
    change_ypos_plantInfo:  ({commit}, payload)  => {
        commit('CHANGE_YPOS_PLANTINFO', payload)
    },
    change_ypos_statInfo: ({commit}, payload) => {
      commit('CHANGE_YPOS_STATINFO', payload)
    },
    clear_farming({commit}){
      commit("CLEAR_FARMING")
    },
    change_dash_state({commit}){
      commit("CHANGE_DASH_STATE")
    },
    change_admin_state({commit}){
      commit("CHANGE_ADMIN_STATE")
    },
    change_batch_state({commit}){
      commit("CHANGE_BATCH_STATE")
    }
  }
})
