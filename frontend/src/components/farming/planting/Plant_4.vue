<template>
<v-container>
 <v-row justify="center"> 
     <h1>Move the plants <br>to the outside</h1>
 </v-row>
  <v-row align-center justify="end" style="padding-left: 50px"> 
      <v-col cols="12" align-self="center" justify-self="end">
          <p>Move each plant one week up</p>
      </v-col>
       <v-col cols="12">
           <img src="../../../assets/harvesting/ModuleWithmissingLetuce.svg">
      </v-col>
  </v-row>

 <v-row align-center justify="end" style="padding-left: 50px"> 
      <v-col cols="12" align-self="center" justify-self="end">
           <p>Put the pot with the <br>seeds into the free <br>spot at week 1</p>
      </v-col>
       <v-col cols="12">
          <img src="../../../assets/harvesting/Module_PlantAtOne.svg">
      </v-col>
  </v-row>
 

    
    </v-container>
</template>

<script>

export default {
    name: "plant_4",
    components:{
   
         
    },
    props:{
       
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