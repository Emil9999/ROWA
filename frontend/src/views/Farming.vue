<template>
 <div id="Farming">
    
      <v-stepper v-model="e1" class="step-number">
      <v-stepper-items>

      <v-stepper-content step="1" >
        <v-row justify="center" align-center>
  <v-col cols="2" style="padding-left: 30px">
      
       
  </v-col>
<v-col align="center"  align-self="center" cols="8">

          <div  style="text-align: center" class="step-header-text">
        <p style="color:#828282">Choose what you would like to do!</p>
          </div>
</v-col>

    <v-col cols="2" style="padding-left: 30px">
          
           <router-link tag="v-btn" :to="{name:'Home'}">
       <v-btn dark fab color="white">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
            </router-link>
  
    </v-col>
        </v-row>
      </v-stepper-content >
 </v-stepper-items>

    <v-stepper-header class="step-number" id="header-steps">
        
      <v-stepper-step :complete=true step="">Name of step 1</v-stepper-step>

      <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 1" step="">Name of step 2</v-stepper-step>

 <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 1" step="">Name of step 3</v-stepper-step>

 <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 2" step="">Name of step 4</v-stepper-step>
      <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 3" step="">Name of step 4</v-stepper-step>
     <v-divider></v-divider>

      <v-stepper-step :complete="e1 > 4" step="">Name of step 4</v-stepper-step>
       
        
    </v-stepper-header>
      </v-stepper>

 
<v-container align="center">
      <v-row justify="center">
        <h1>Harvest a plant</h1>
      </v-row>

      <v-row justify="center">
        <p>Go on a journey to become a famer.Harvest a plant now, for your daily dose of fresh greens.</p>
      </v-row>
      <v-row justify="center" class="info-box" margin="auto" padding="auto">
        <InfoBoxPlants heading="Harvestable" :plants="harvestable_plants"></InfoBoxPlants>
      </v-row>

      <v-row justify="center" style="margin-top: 40px">
        <router-link :to="{name:'Harvest'}">
          <v-btn id="button" rounded color="accent" height="75" width="360">
            Start Harvesting
            <v-icon right dark>mdi-arrow-right</v-icon>
          </v-btn>
        </router-link>
      </v-row>

      <v-row justify="center">
        <h1>Plant a new Plant</h1>
      </v-row>

      <v-row justify="center">
        <p>Go on a journey to become a famer.Plant a new seed now and enjoy fresh greens in the future.</p>
      </v-row>
      <v-row justify="center" class="info-box" margin="auto" padding="auto">
        <InfoBoxPlants heading="Plantable" :plants="plantable_plants"></InfoBoxPlants>
      </v-row>

      <v-row justify="center" style="margin-top: 40px">
        <router-link :to="{name:'Plant'}">
        <v-btn id="button" rounded color="accent" height="75" width="360">
          Start Planting
          <v-icon right dark>mdi-arrow-right</v-icon>
        </v-btn>
        </router-link>
      </v-row>
    
    </v-container>
    </div>
</template>

<script>
import axios from "axios";

import InfoBoxPlants from "../components/home/InfoBoxPlants";
export default {
  name: "Farming",
  components: {
   
    InfoBoxPlants
  },
  data() {
    return {
      step: 0,
      harvestable_plants: null,
      plantable_plants: null
    };
  },
  methods: {
    getHarvestablePlants: function() {
      axios
        .get("http://127.0.0.1:3000/dashboard/harvestable-plants")
        .then(result => {
          this.harvestable_plants = result.data;
          this.plantable_plants = result.data;
          console.log(this.plantable_plants)
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    //Get request to get all plants in farm + harvestable plants
    this.getHarvestablePlants();
  }
};
</script>

<style scoped>
#button {
  font-weight: bold;
  margin: 40px;
  font-family: Montserrat;
  font-size: 24px;
}
h1 {
  color: #789659;
}

p {
  text-align: center;
  width: 410px;
  color: #828282;
  padding-top: 10px;
  margin: 0 !important;
}
.info-box {
  background: #789659;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 100px 0 100px;
}
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
</style>