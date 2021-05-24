<template>
  <div id="RealityCheck">
    <div v-if="!moduleSelected">
    <AdminTopRow v-bind:headtext="'Reality Check'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
    <v-row justify="center" style="margin: 20px 0 0px 0">
    <h1 style="color:#789659">Click on the module that appears to have wrong plant data</h1>
    </v-row>
     <Cattree class="cattree"  v-on:moduleClicked="onChildClick"></Cattree>
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
 <div>
    <Module class="position-info" :id="this.id" :reverse="this.reverse" ></Module>
  </div>
 <div style="max-height:630px; overflow:auto;">
  <div v-for="key in unrevpositions" :key="key">
                <div v-if="module_plants.pos[key].age!=0">
                          <v-row class="info-box" justify="center" align-content="center" align="center"> 
     
       <v-col cols="2" align-self="center" align="center"> <h4>Type: </h4><br> <h5>{{module_plants.type.replace(/\b\w/g, l => l.toUpperCase())}}</h5>  </v-col>
       <v-col cols="2" align-self="center" align="center"> <h4>Position: </h4>  <h5>{{key}} </h5></v-col>
        <v-col cols="3" > <img :src="getImgUrl(module_plants.type)" alt="" width="120px" height="auto"> </v-col>
        <v-col cols="3" > <v-select
                  v-model="module_plants.pos[key].age"
                  :items="weeks"
                  menu-props="auto"
                  item-text="week"
                  item-value="days"
                  :label="'Week '+Math.round((module_plants.pos[key].age)/7)"
               
                ></v-select>
                </v-col>
              <v-col cols="2"><v-btn id="button" rounded color="accent" height="25" width="25" @click="removePlant(key)"> 
                  <v-icon color="red">mdi-close</v-icon>
                </v-btn></v-col>
     
   </v-row>
                </div>
                <div v-else>

                  <v-row class="info-box" justify="center" align-content="center" align="center"> 
     
       <v-col cols="2" align-self="center" align="center"> <h4>Type: </h4><br> <h5>Empty</h5>  </v-col>
       <v-col cols="2" align-self="center" align="center"> <h4>Position: </h4>  <h5>{{key}} </h5></v-col>
        <v-col cols="3">  <img src="../assets/admin/blackX.png" alt="" width="auto" height="70px"> </v-col>
       <v-col cols="2"> </v-col>
        <v-col cols="3"> <v-select
                  v-model="module_plants.pos[key].age"
                  :items="weeks"
                  item-text="week"
                  item-value="days"
                  menu-props="auto"
                  hide-details
                
                  :label="'Weeks'+(module_plants.pos[key].age)%7"
                ></v-select>
                </v-col>
     
   </v-row>
            
                </div>
  </div>
 </div>

             <v-row justify="center" style="margin-top: 40px"> 
                <v-btn id="button" rounded color="accent" height="50" width="360" @click="sendChangesRealityCheck()">Save 
                  <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
                </v-btn>
    </v-row>
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
  props: ["id"],
  components: {
    AdminTopRow,
    Cattree,
    Module,
    
  },

  data() {
    return {
      appendSlot: true,
      moduleClicked: 0,
      moduleSelected: false,
      weeks: [{week: "Empty",days:0},
              {week: "Week 0",days:2},
              {week: "Week 1",days:7},
              {week: "Week 2",days:14},
              {week: "Week 3",days:21},
              {week: "Week 4",days:28},
              {week: "Week 5",days:35},
              {week: "Week 6",days:42},
              {week: "Week 6+",days:50}
              ],


    };
  },
  computed: {
     ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.id]
            },
            reverse: function (){
              if(this.id%2==0){
                return  true
              }
              else{
                return false
              }
            },
            positions: function () {
                if (this.reverse){
                    return Object.keys(this.module_plants.pos)
                }
                else {
                    return Object.keys(this.module_plants.pos).reverse()
                    
                }
              
            },
            unrevpositions: function () {
                    return Object.keys(this.module_plants.pos)
                }

  },
  methods: {
     ...mapGetters(["get_module"]),
     onChildClick (value) {
      this.id = value
      this.moduleSelected = true
      
    },
    removePlant(modulenum){
      this.module_plants.pos[modulenum].age = 0
      this.module_plants.pos[modulenum].harvestable = null

    },
      getImgUrl(pic) {
      return require("@/assets/harvesting/plants/" + pic.replace(/\b\w/g, l => l.toUpperCase()) + ".png");
    },
      makeOutput(){
        var postions = new Array;
        var x;
        for (x in this.module_plants.pos){
            postions.push(this.module_plants.pos[x].age)
        }
        var output={type: this.module_plants.type, modulenum: this.id, age: postions}
        console.log(output)
        return (output)
      },
      /*getKnownTypes: function() {
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
    },*/
    sendChangesRealityCheck: function(){
      axios
      .post( "http://127.0.0.1:3000/admin/reality-check",
          this.makeOutput(),
          "content-type: application/json"
        )
      .then()
      .catch(error => {
          console.log(error);
      });
    },/*
      populateModule(moduleNumber){
                axios.get("http://127.0.0.1:3000/dashboard/cattree/"+moduleNumber.toString())
                    .then(result => {
                        
                        result.data.moduleNumber = moduleNumber.toString()
                        this.$store.commit("FarmUpdate", result.data)
                    })
                    .catch(error => {
                        console.log(error)
                    })
                },*/
                
            
  
  },
   updated(){
          for (var x in this.module_plants.pos){
        if(this.module_plants.pos[x].age==0 && this.module_plants.pos[x].harvestable!=null){
          this.module_plants.pos[x].age=1
        }
      }
   },
   created() {
    /* for(var i = 1;i<7;i++){
                this.populateModule(i)
            }*/
    
   
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

  .position-info{
    position:absolute;
    top: 0px;
    left: 210px;
    transform: scale(1.5);
}

 .cattree{
    position:absolute;
    top: 170px;
    left: 100px;
    transform: scale(1.3);
}
</style>