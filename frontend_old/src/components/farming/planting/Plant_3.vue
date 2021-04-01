<template>
<v-container style="margin: 0px 0 0px 0">
<v-row justify="center" style="margin: -20px 0 0px 0">
    <h1  style="color:#789659">Put the seed into the pot</h1>
</v-row>

<v-row justify="center" >
<v-spacer></v-spacer>
    <v-col  cols="2"   justify-end>
        
     <img src="../../../assets/harvesting/rockwool.svg" height="auto" style="margin:0 0px 0 0; padding:0 0px 0 0;">
    </v-col>
    <v-col cols="6">
    <p>Get the seed from  <br> the seed dispenser.</p>
    </v-col>
</v-row>

    
  
<v-row justify="center">
    <v-spacer></v-spacer>
    <v-col  cols="5"   justify-end>
     <img src="../../../assets/harvesting/Pot.svg" height="auto">
    </v-col>
    <v-col cols="6">
    <p>Take a pot from <br>the pot dispenser<br>(underneath the seeds).</p>
    </v-col>
     
</v-row>
<v-row justify="center" style="margin: 30px 0 0 0">
         <v-col cols="5">
    <div style="padding:0px 0px 20px 0; margin: 0 0px 0 -30px;">
    <img src="../../../assets/harvesting/rockwool.svg" height="auto">
   <img src="../../../assets/harvesting/arrow.svg" height="auto">
    </div>
    <img src="../../../assets/harvesting/Pot.svg" height="auto">
     </v-col>
    <v-col cols="6">
    <p>Put the seed into the pot and press it firmly down until you reach the bottom.</p>
    </v-col>

</v-row>


 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt=""  width="120px" height="auto">> </v-col>
       
   </v-row>

    </v-container>
</template>

<script>



export default {
    name: "plant_3",
    components:{
        
    },
    props:{
        selectedPlant:String
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
</style>