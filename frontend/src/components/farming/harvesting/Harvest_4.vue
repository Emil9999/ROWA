<template>
<v-container>
 <v-row justify="center"> 
     <h1>Seperate the plot from the plant</h1>
 </v-row>
  <v-row align-center justify="start" style="padding-left: 50px; margin-top: 20px;"> 
   
      <v-col cols="6" align-self="center" justify-self="end">
          <img src="../../../assets/harvesting/lettuce.png">
      </v-col>
       <v-col cols="6">
           <img src="../../../assets/harvesting/Pot.svg">
      </v-col>
  </v-row>

 <v-row align-center style="padding-left: 50px"> 
      <v-col cols="6" align-self="center">
           <p>Put the pot with the roots into the disposal-box<br> (in the lower right-hand corner). The pot will be washed and reused in the future.</p>
      </v-col>
       <v-col cols="6">
          <img src="../../../assets/harvesting/InsertionInstructions.svg">
      </v-col>
  </v-row>
 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
       
   </v-row>

    
    </v-container>
</template>

<script>

export default {
    name: "harvest_4",
    components:{
   
         
    },
    props:{
        selectedPlant: String,
        posandModule: Array
    },
    methods:{
        getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic+".png")
            },
    },

    
}
</script>


<style scoped>
.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 100px 0 100px;

}
h3{
      
font-family: Montserrat;
font-style: normal;
font-weight: 600;
font-size: 18px;
line-height: 22px;
color:var(--v-primary-base)
}
#button{
 font-weight: bold;
 margin: 40px;
 font-family: Montserrat;
 font-size: 24px
}
h1{
    
      
font-family: Montserrat;
font-style: normal;
font-weight: 600;
text-align: center;

color:var(--v-primary-base)
}
</style>