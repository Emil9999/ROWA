<template>
<v-container>
<v-row justify="center">
    <h1  style="color:#789659">Remove the plant with the pot now</h1>
</v-row>
<v-row justify="center">
    <p>this shows a module telling you the plant position: {{posandModule[1]}}</p>
</v-row>

    
    <posInstructions></posInstructions>

<v-row justify="center">
    <p>Move the pot slightly to the widened area for easy removal</p>
</v-row>
<v-row justify="center">
    
    <img src="../../../assets/harvesting/pickoutInstructions.svg" height="120px">
</v-row>


 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img src="../../../assets/logo.svg" alt="" width="80px" height="auto"> </v-col>
       
   </v-row>

    </v-container>
</template>

<script>


import posInstructions  from "../harvesting/posInstructions"
export default {
    name: "harvest_3",
    components:{
         posInstructions
    },
    props:{
        selectedPlant: String,
        posandModule: Array
    },

    
}
</script>


<style scoped>

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 20px 90px 0 90px;
}

#button{
 font-weight: bold;
 margin: 40px;
 font-family: Montserrat;
 font-size: 24px
}
</style>