<template>
<v-container style="margin: 0px 0 0px 0">
<v-row justify="center">
    <h1  style="color:#789659">You are Planting:</h1>
</v-row>
          
<v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> <p>Module: {{module}}</p></v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
        
       
   </v-row>
    
<v-row justify="center">
    <h1  style="color:#789659">Is this your first time?</h1>
</v-row>
<v-row justify="center">
<v-btn id="button" style="margin:5px;" min-height="75px" min-width="400px" color="primary" x-large rounded @click="gotoInstructions(true)">Watch Instructions<v-icon size="1.5rem">mdi-arrow-right</v-icon></v-btn>
</v-row>

<v-row class="harvest-box" justify="center"> 
    <v-row justify="center">
     <h1  style="color:#789659">Already know how to plant?</h1>
   </v-row>
   <v-row justify="center">
           <p>By clicking you skip the Instructions Video</p>
   </v-row>
   <v-row justify="center">
<v-btn id="button" class="green--text" style="margin:5px;" fab min-height="75px" min-width="400px" color="white" elevation="22" x-large rounded @click="gotoInstructions(false)">Start Planting</v-btn>
</v-row>
   </v-row>


    </v-container>
</template>

<script>



export default {
    name: "plant_3",
    components:{
        
    },
    props:{
        selectedPlant:String,
        module: Number,
    },
    methods:{
         gotoInstructions(knowUser){
            this.$emit("gotoInstructions", knowUser)
        },
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
  margin: 40px 100px 80px 100px;

}

.harvest-box {
  background: #ffffff;
  border-radius: 20px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 60px 20px 30px 20px;
  padding: 50px;

}

.position-info{
    position:absolute;
    top: 0px;
    left: 0x;
    transform: scale(1.75);
}
.position-info-box{
    position:absolute;
    top: 100px;
    left: 230px;
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