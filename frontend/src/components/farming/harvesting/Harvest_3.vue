<template>
<v-container>
<v-row justify="center">
    <h1  style="color:#789659">Follow the instructions to take your plant from the module</h1>
</v-row>

<v-row justify="center" style="margin: -20px 0 45px 0 ;">
    <Module :id="this.module" :reverse="false" :pos="this.pos"></Module>
    
</v-row>

    


<v-row justify="center">
    <ol><li> Move the plant with the pot to the widened area of the module for easy removal. </li> 
        <li> After lifting the plant wait for the water to stop dripping from the roots </li>
      <li>Remove the plant from the system</li>
        </ol>
</v-row>
<v-row justify="center">
    
    <img :src="getImgUrlForRemoveInstruc()" height="120px">
</v-row>


 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> </v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
       
   </v-row>

    </v-container>
</template>

<script>

import Module from "@/components/farming/harvesting/ModulePostitionIndi.vue"

export default {
    
    name: "harvest_3",

    components:{
        Module
 
    },
    props:{
        selectedPlant: String,
        pos: Number,
        module: Number
    },
    methods:{
         getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic+".png")
            },
        getImgUrlForPos() {
            if(this.module%2 == 0){
                return require('@/assets/harvesting/WhereToHarvest/ModulewithLettuceRightPos'+this.pos+".svg")
            } else {
                return require('@/assets/harvesting/WhereToHarvest/ModulewithLettuceLeftPos'+this.pos+".svg") 
            }
            },
        getImgUrlForRemoveInstruc(){
             if(this.module%2 == 0){
                return require('@/assets/harvesting/pickoutInstructionsright.svg')
            } else {
                return require('@/assets/harvesting/pickoutInstructionsleft.svg') 
            }
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