<template>
  <div id="RealityCheck">
    <div v-if="!moduleSelected">
    <AdminTopRow v-bind:headtext="'Reality Check'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
     <Cattree  v-on:moduleClicked="onChildClick"></Cattree>
     <h1>{{this.moduleClicked}}</h1>
    </div>
    <div v-else>
      <v-row align-center>
        <v-col cols="2" style="padding-left: 30px">
            <v-btn dark fab color="white" v-on:click="moduleSelected = false" >
                <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
        </v-col>
        <v-col align="center"  align-self="center" cols="8">
             <h3 style="color:#828282">Reality Check</h3>
        </v-col>
        <v-col cols="2" id="logo-text" style="text-align: right; padding-right: 30px">
            <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
        </v-col>
    </v-row>
 <div class="module" style="width: 600px; object-fit: fill; right: 0; display:flex">
    <Module :id="this.id" :reverse="this.reverse" style=""></Module>
  </div>
  <div v-for="key in positions" :key="key">
                <div v-if="module_plants.pos[key].age!=0">
                          <v-row class="info-box" justify="center" align-content="center" align="center"> 
     
       <v-col align-self="center" align="center"> <h4>Plant-Type: </h4><br> <h5>{{module_plants.type}}</h5>  </v-col>
       <v-col align-self="center" align="center"> <h4>Position: </h4>  <h5>{{key}} </h5></v-col>
        <v-col > <img :src="getImgUrl(module_plants.type)" alt="" width="120px" height="auto"> </v-col>
        <v-col > <v-select
                  v-model="module_plants.pos[key].age"
                  :items="weeks"
                  menu-props="auto"
                  item-text="week"
                  item-value="days"
                  :label="'Week '+Math.round((module_plants.pos[key].age)/7)"
                  hide-details
                  single-line
                ></v-select></v-col>
     
   </v-row>
                </div>
                <div v-else>

                  <v-row class="info-box" justify="center" align-content="center" align="center"> 
     
       <v-col align-self="center" align="center"> <h4>Plant-Type: </h4><br> <h5>{{module_plants.type}}</h5>  </v-col>
       <v-col align-self="center" align="center"> <h4>Position: </h4>  <h5>{{key}} </h5></v-col>
        <v-col > <img src="../assets/admin/blackX.png" alt="" width="auto" height="70px"> </v-col>
        <v-col > <v-select
                  v-model="module_plants.pos[key].age"
                  :items="weeks"
                  item-text="week"
                  item-value="days"
                  menu-props="auto"
                  hide-details
                  single-line
                  :label="'Weeks'+(module_plants.pos[key].age)%7"
                ></v-select></v-col>
     
   </v-row>
            
                </div>
            </div>
    </div>
  </div>
</template>


<script>
import AdminTopRow from "@/components/admin/AdminTopRow.vue";
import Cattree from "@/components/home/Farm/CatTree.vue";
import {mapState, mapGetters} from "vuex"
import axios from "axios";
import Module from "@/components/home/Farm/Module.vue";

export default {
  name: "RealityCheck",
  props: ["id", "reverse: true"],
  components: {
    AdminTopRow,
    Cattree,
    Module,
    
  },

  data() {
    return {
      moduleClicked: 0,
      moduleSelected: false,
      weeks: [{week: "Empty",days:0},
              {week: "Week 1",days:7},
              {week: "Week 2",days:14},
              {week: "Week 3",days:21},
              {week: "Week 4",days:28},
              {week: "Week 5",days:35},
              {week: "Week 6",days:42},
              {week: "Week 6+",days:50}
              ],
      types: []

    };
  },
  computed: {
     ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.id]
            },
            positions: function () {
                if (this.reverse){
                    return Object.keys(this.module_plants.pos)
                }
                else {
                    return Object.keys(this.module_plants.pos).reverse()
                }
            }

  },
  methods: {
     ...mapGetters(["get_module"]),
     onChildClick (value) {
      this.id = value
      this.moduleSelected = true
      
    },
    addPlant(modulenum){
      this.module_plants.pos[modulenum].age = 1
      this.module_plants.pos[modulenum].harvestable = false

    },
      getImgUrl(pic) {
      return require("@/assets/harvesting/plants/" + pic.replace(/\b\w/g, l => l.toUpperCase()) + ".png");
    },
       getKnownTypes: function() {
      axios
        .get("http://127.0.0.1:3000/adminSettings/get-knowntypes")
        .then(result => {
          result.data.forEach(element => {
            this.types.push(element.type);
          });
        })
        .catch(error => {
          console.log(error);
        });
    },
  },
   created() {
    this.getKnownTypes();
  }
};
</script>

<style scoped>
h1,
h2 {
  font-weight: normal;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: #42b983;
}

span {
  color: var(--v-primary-base);
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
}

.no-hover:hover {
  background-color: transparent;
  text-decoration: none;
}

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}

  #logo-text{
        font-family: Montserrat, sans-serif;
        font-style: normal;
        font-weight: 600;
        font-size: 36px;
        color: var(--v-primary-base)
    }
</style>