<template>
<v-container justify="center" >
   <!--<FarmingTopRow v-bind:step="1" v-bind:headtext="'Select a plant-type'" v-bind:previosPage="'Farming'"></FarmingTopRow>-->
    <div  v-for="item in available_plants" :key="item.plant_type">
       
   <v-row class="info-box" justify="center" v-on:click="sendSelectedPlant(item.plant_type)"> 
     
       <v-col align-self="center" align="center"> <h3>{{item.plant_type}} </h3> </v-col>
        <v-col> <img src="../../../assets/logo.svg" alt="" width="120px" height="auto"> </v-col>
     
   </v-row>
   
    </div>
    </v-container>
   
</template>

<script>
import axios from "axios"

export default {
    name: "plant_1",
    components:{
   
    },

    data(){
        return{
           
        available_plants:null
    }
    },
 
        
    methods:{
        sendSelectedPlant(plant){
            this.$emit('sendSelectedPlant',  plant)
        },
         getHarvestablePlants: function () {
                axios.get("http://127.0.0.1:3000/dashboard/plantable-plants")
                    .then(result => {
                        this.available_plants = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
    },
    created() {
            //Get request to get all plants in farm + harvestable plants
            this.getHarvestablePlants()}
}    
    

</script>

<style scoped>
.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 100px 0 100px;
}
</style>