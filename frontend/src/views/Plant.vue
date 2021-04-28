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
       
       <v-btn dark fab color="white"  :to="{name:'Home'}">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
       
  
    </v-col>
        </v-row>
      </v-stepper-content >

      <v-stepper-content step="2" class="step-header-text">
         <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
        
       <v-btn dark fab color="white" @click="e1 -=1" v-on:click="abortBlinking()">
           <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
       
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Find your module</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
         
       <v-btn dark fab color="white"  v-on:click="abortBlinkingHome()">
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
        <p style="color:#828282">Prepare</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
         
       <v-btn dark fab color="white" v-on:click="abortBlinkingHome()">
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
        <p style="color:#828282">Plant</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
         
       <v-btn dark fab color="white" v-on:click="abortBlinkingHome()"> 
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
         <Plant_1 v-on:sendSelectedPlant="nextStepAndSaveplant($event)"></Plant_1>
      </v-stepper-content>


      <v-stepper-content step="2" style="margin:-25px 0px -25px 0px;">  
       <Plant_2 v-bind:selectedPlant="this.selectedPlantType" v-bind:moduleNumber="this.moduleNum"></Plant_2>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="50" width="360" @click="e1 = 3">
         See Instructions
         <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


      <v-stepper-content step="3">
       <Plant_3 v-bind:selectedPlant="this.selectedPlantType"></Plant_3>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="75" width="360" @click="e1 = 4">
          Next
          <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


       <v-stepper-content step="4">
         <Plant_4 v-bind:selectedPlant="this.selectedPlantType"
                  v-bind:module="this.moduleNum"></Plant_4>
       <v-row justify="center">
        <v-btn id="button" rounded color="accent" height="75" width="360" @click="e1 = 5" v-on:click="advanceTimer()">
         I Planted 
         <v-icon>mdi-arrow-right</v-icon>
        </v-btn>
       </v-row>
      </v-stepper-content>


       <v-stepper-content step="5">
       
         <v-row justify="center" align-self="end"  v-on:click="sendPlantedPlant()">
          
      <img src="../assets/harvesting/confetti.svg" alt="">
         </v-row>
         <v-row justify="center" style="margin:20px; text:block" color="primary">
           <h1>You have successfully <br>planted</h1>
         </v-row>

      </v-stepper-content>


    </v-stepper-items>
  </v-stepper>
</template>

<script>
import axios from "axios"
import Plant_1 from "../components/farming/planting/Plant_1"
import Plant_2 from "../components/farming/planting/Plant_2"
import Plant_3 from "../components/farming/planting/Plant_3"
import Plant_4 from "../components/farming/planting/Plant_4"
 import {mapState} from "vuex"
 
  export default {
      name:"Harvest",
      components:{
          Plant_1,
          Plant_2,
          Plant_3,
          Plant_4
      },
    data () {
      return {
        e1: 1,
        selectedPlantType: "Basil",
          autoAdvanceTimer: null,
        moduleNum: 0,
        
      }
    },
    computed:{
        ...mapState(["to_farm"]),
    },
    
    methods:{
      getPositonAndModuleOfPlant:function(){
        axios.
        get("http://127.0.0.1:3000/plant/get-position",
        {params: {
                        plant_type: this.selectedPlantType
                    }},
          "content-type: application/json")
        .then(result => {
          this.moduleNum = result.data
         
          

        })
        .catch(error => {
          console.log(error);
        });
      },

      nextStepAndSaveplant: function(selectedPlanttype){
          this.selectedPlantType = selectedPlanttype;
          this.e1 +=1;
          this.getPositonAndModuleOfPlant();
      },

      sendPlantedPlant:function(){
        axios.post("http://127.0.0.1:3000/plant/finish",
        {planted_module:this.moduleNum },
          "content-type: application/json")
        .then()
        this.$store.dispatch('clear_farming')
         this.$router.push('/')
        .catch(error => {
          console.log(error);
        });
      },
      abortBlinking:function(){
        axios.get("http://127.0.0.1:3000/light/"+this.moduleNum+"/breath-off")
        .then()
        this.$store.dispatch('clear_farming')
        .catch(error => {
          console.log(error);
        });
      },
       abortBlinkingHome:function(){
        axios.get("http://127.0.0.1:3000/light/"+this.moduleNum+"/breath-off")
        .then()
        this.$store.dispatch('clear_farming')
        this.$router.push('/')
        .catch(error => {
          console.log(error);
        });
      },
     advanceTimer () {
            this.autoAdvanceTimer = setInterval(() => {
                this.sendPlantedPlant()
            }, 3000)
        },
            
    },
     beforeDestroy () {
	clearInterval(this.autoAdvanceTimer)
},

      created() {
        if(this.to_farm.module != 0){

          this.moduleNum = this.to_farm.module
          this.selectedPlantType = this.to_farm.plant_type
          this.e1 = 2
      }
        
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

#button {
  font-weight: bold;
  margin: 40px;
  font-family: Montserrat;
  font-size: 24px;
  font-family: Montserrat;
font-style: normal;

}

</style>