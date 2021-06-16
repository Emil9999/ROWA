<template>
<v-container>
<v-overlay
          color="white"
          absolute="true"
          opacity="1"
          :value="VideoInst"
        > 
          
          <v-row justify="center" margin="100px"> 
                 <video autoplay="true" height="700" width="400"  @ended="VideoInst = false">
                <source :src="require(`/Users/hannesostwald/Downloads/1-Susteyn 3min Elevator Pitch.mp4`)" type="video/mp4">
    </video>
          </v-row>
          <v-row justify="center"> 
            <v-btn
            color="success"
            @click="VideoInst = false"
          >
           Skip
          </v-btn>
                 </v-row>

                
           </v-overlay>

 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> <p>Module: {{module}}</p></v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
       
   </v-row>

 <v-row justify="center"> 
     <h1>Harvest a Lettuce</h1>
 </v-row>
  <v-row align-center justify="start" style="padding-left: 50px; margin-top: 20px;"> 
   <ol>
       <li>Take a Drip-Container</li>
       <li>Remove the Plant from the Module</li>
       <li>Seperate the Plant from the Pot</li>
       <li>Return Drip-Container and Pot</li>
   </ol>
    
  </v-row>

<v-row justify="center">
<v-btn id="button" style="margin:5px;" min-height="75px" min-width="400px" color="primary" x-large rounded @click="restartInstruc()">Watch Instructions</v-btn>
</v-row>

    
    </v-container>
</template>

<script>

//import VideWin from '@/components/farming/VideoPlayer.vue'

export default {
      data () {
        return {
        VideoInst: true
    };
      },
    name: "harvest_4",
    components:{
        //VideWin
   
         
    },
    props:{
        selectedPlant: String,
        module: Number
    },
    methods:{
        getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic+".png")
            },
        restartInstruc() {
            this.VideoInst = true
        }
    },

    
}
</script>


<style scoped>
.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 100px 40px 100px;

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
li{
    margin: 40px;
    line-height: 22px;
    font-family: Montserrat;
    font-size: 24px;
    color:var(--v-primary-base)
}
h1{
    
      
font-family: Montserrat;
font-style: normal;
font-weight: 600;
text-align: center;

color:var(--v-primary-base)
}
</style>