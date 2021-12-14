<template>
<v-container style="height:1100px">
<v-overlay
          color="white"
          absolute="true"
          opacity="1"
          :value="VideoInst"
          
        > 
          
          <v-row justify="center" margin="0px"> 
                 <video ref="VidIn" autoplay="true" muted  height="900" width="auto"  @ended="videoInst = false"
                 >
                <source :src="VideoUrl" type="video/mp4">
    </video>
          </v-row>
         
                  <v-row justify="center"> 
          <v-btn id="button" style="margin:10px;" underline text min-height="75px" min-width="400px" color="error"  rounded  @click="videoInst = false">Skip Instructions</v-btn>
                 </v-row>
  
           </v-overlay>

 <v-row class="info-box" justify="center"> 
      
       <v-col align-self="center" align="center"> <h3 v-text="selectedPlant"></h3> <p>Module: {{module}}</p></v-col>
        <v-col> <img :src="getImgUrl(this.selectedPlant)" alt="" width="120px" height="auto"> </v-col>
       
   </v-row>
 <v-row class="harvest-box" justify="center"> 
 <v-row justify="center"> 
     <h1>Harvesting Instructions:</h1>
 </v-row>
  <v-row align-center justify="start" style="padding-left: 50px; margin-top: 20px;"> 
   <ol>
       <li>Get a Drip-Pot(for Step 3)</li>
       <li>Move the plant to the outside</li>
       <li>Take Plant &amp; Place it into the Drip-Pot</li>
       <li>Return Drip-Container and Pot</li>
   </ol>
    
  </v-row>
   <v-row justify="center">
                    <v-btn id="button" rounded color="primary" height="75" width="400" @click="e1 = 4" v-on:click="goToFinal()">
                        I Harvested a Lettuce
                        <v-icon>mdi-arrow-right</v-icon>
                    </v-btn>
                </v-row> </v-row>

<v-row justify="center">
<v-btn id="button" style="margin:5px;" underline text min-height="75px" min-width="400px" color="primary"  rounded @click="restartInstruc()">Watch Instructions again</v-btn>
</v-row>

    
    </v-container>
</template>

<script>



export default {
      data () {
        return {
          

    };
      },
      
    name: "harvest_4",
    components:{
   
         
    },
    props:{
        videoInst: Boolean,
        selectedPlant: String,
        module: Number
    },
    computed:{
      VideoUrl: function(){
          if(this.selectedPlant != "Herb"){
            return require(`@/assets/videos/HarvestInst.mp4`)
          }else{
            return require(`@/assets/videos/HarvestInst.mp4`)
          }},
          VideoInst: function (){
            return this.videoInst
          }
          },
     created: function() {
       this.videoInst = false;
     },
    methods:{
        restart(){
          this.$refs.VidIn.play();
        },
        pause(){
           this.$refs.VidIn.pause();
        },
        getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic+".png")
            },
        restartInstruc() {
            this.videoInst = true
        },
         setUpInstruc() {
            this.videoInst = false
        },
        goToFinal(){
          this.$emit("goToFinal")
        }
    }, 

    
}
</script>


<style scoped>
.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 0px 100px 30px 100px;

}
.harvest-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 10px 10px 0 10px;
  padding: 20px;
 

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
 margin: 30px;
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