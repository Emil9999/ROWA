<template>
  <v-stepper v-model="e1" class="step-number">
      <v-stepper-items>

      <v-stepper-content step="1" >
        <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
      
        <v-btn dark fab color="white" :to="{name:'Farming'}">
           <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Select plant-type</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
         
       <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
 
  
    </v-col>
        </v-row>
      </v-stepper-content >

      <v-stepper-content step="2" class="step-header-text">
         <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
        
       <v-btn dark fab color="white" @click="e1 -=1">
           <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
       
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Find your plant</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
      
       <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
        
  
    </v-col>
        </v-row>
      </v-stepper-content>

      <v-stepper-content step="3" class="step-header-text">
          <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
        
       <v-btn dark fab color="white" @click="e1 -=1">
           <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
       
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Harvest</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
      
       <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
     
  
    </v-col>
        </v-row>
      </v-stepper-content>

       <v-stepper-content step="4" class="step-header-text">
          <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
        
       <v-btn dark fab color="white" @click="e1 -=1">
           <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
       
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Finishing Up</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
        
       <v-btn dark fab color="white" :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
      
  
    </v-col>
        </v-row>
      </v-stepper-content>

       <v-stepper-content step="5" class="step-header-text">
          <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
   
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Success</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
           
    </v-col>
        </v-row>
      </v-stepper-content>

      </v-stepper-items>

    <v-stepper-header class="step-number" id="header-steps">
        
      <v-stepper-step :complete=true step="">Name of step 1</v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete=true step="">Name of step 2</v-stepper-step>

 <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 1" step="">Name of step 3</v-stepper-step>

 <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 2" step="">Name of step 4</v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 3" step="">Name of step 4</v-stepper-step>
     <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 4" step="">Name of step 4</v-stepper-step>
       
        
    </v-stepper-header>

    <v-stepper-items color="secondary">
      <v-stepper-content step="1" color="secondary">
         <Harvest_1 v-on:sendSelectedPlant="nextStepAndSaveplant($event)"></Harvest_1>
      </v-stepper-content>


      <v-stepper-content step="2">
       <Harvest_2 v-bind:selectedPlant="this.selectedPlantType" v-bind:moduleNumber="this.moduleNum"></Harvest_2>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="75" width="360" @click="e1 = 3">
         Got it
         <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


      <v-stepper-content step="3">
       <Harvest_3 v-bind:selectedPlant="this.selectedPlantType"
        v-bind:posandModule="this.position"></Harvest_3>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="75" width="360" @click="e1 = 4">
          Next
          <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


       <v-stepper-content step="4">
         <Harvest_4 v-bind:selectedPlant="this.selectedPlantType"></Harvest_4>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="75" width="360" @click="e1 = 5">
         I Harvested 
         <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


       <v-stepper-content step="5">
         <router-link :to="{name:'Home'}" v-on:click="sendPlantedPlant()">
         <v-row justify="center" align-self="end">
           
      <img src="../assets/harvesting/confetti.svg" alt="">
         </v-row>
         <v-row justify="center" style="margin:20px; text:block" color="primary">
           <h1>You have successfully <br>harvested</h1>
         </v-row>
         </router-link>
      </v-stepper-content>


    </v-stepper-items>
  </v-stepper>
</template>

<script>
import axios from "axios"
import Harvest_1 from "../components/farming/harvesting/Harvest_1"
import Harvest_2 from "../components/farming/harvesting/Harvest_2"
import Harvest_3 from "../components/farming/harvesting/Harvest_3"
import Harvest_4 from "../components/farming/harvesting/Harvest_4"
  export default {
      name:"Harvest",
      components:{
          Harvest_1,
          Harvest_2,
          Harvest_3,
          Harvest_4
      },
    data () {
      return {
        e1: 1,
        selectedPlantType: "Basil",
      
        
        moduleNum: 2,
        position: 6
      }
    },
    methods:{
      getPositonAndModuleOfPlant:function(){
        axios.
        post("http://127.0.0.1:3000/harvest/get-plant",
        { plantType: this.selectedPlantType},
          "content-type: application/json")
        .then(result => {
          this.moduleNum = result.data[0],
          this.position = result.data[1]
          

        })
        .catch(error => {
          console.log(error);
        });
      },

      nextStepAndSaveplant: function(selectedPlanttype){
          this.selectedPlantType = selectedPlanttype;
          this.e1 +=1;
          this.getPositonAndModuleOfPlant();
          console.log(this.position)
      },

      sendPlantedPlant:function(){
        axios.
        post("http://127.0.0.1:3000/harvest/harvestdone",
        { plant_Position: this.position, module_Position: this.moduleNum },
          "content-type: application/json")
        .then()
        .catch(error => {
          console.log(error);
        });
      },
    }
  }
</script>

<style scoped>
#header-steps{
padding: 0px 130px 0px 130px;
margin-top:-40px;
} 
 .step-number{
  background-color:var(--v-secondary-base) !important;
  border-color: var(--v-secondary-base) !important;
  box-shadow: none !important;
 
  
}

.step-header-text{
  
    color: var(--v-primary-base) !important;
    
}
p{
  font-family: Montserrat;
font-style: normal;
font-weight: 500;
font-size: 24px;
line-height: 29px;
}
#button{
 font-weight: bold;
 margin: 40px;
 font-family: Montserrat;
 font-size: 24px; 
}
</style>