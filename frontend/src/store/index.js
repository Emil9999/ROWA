import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)



class Module {
  constructor(type){
    this.type = type;
    this.pos = [ {
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
  ]
  }
}


const farmInfo = {
  state: () => ({ 
    1: new Module("lettuce"),
    2: new Module("lettuce"),
    3: new Module("lettuce"),
    4: new Module("lettuce"),
    5: new Module("lettuce"),
    6: new Module("lettuce")

  })
  
}




export default new Vuex.Store({
  state: {
    farmActive:true
  },
  mutations: {
  },
  actions: {
  },
  modules: { //Vuex modules, not farm modules
    farmInfo: farmInfo
  }
})
