<template>
    <FarmTransition :y-positions="yPositions">
        <div>
            <v-row justify="space-around"  aling-self="center">
                <v-col cols="1">
                  <v-btn dark fab small color="white" @click="upYpos()">
            <v-icon color="primary">mdi-arrow-up</v-icon>
          </v-btn></v-col> 
              <v-col cols="auto" justify="center"></v-col>
           <v-col cols="1">      <v-btn dark fab small color="white" @click="setYpos()">
            <v-icon color="primary">mdi-close</v-icon>
          </v-btn></v-col>
            </v-row>
            <v-row justify="center">  <h2>{{statText[upperCaseStat].In.title}}</h2></v-row>
            <v-row justify="center">
                <p>{{statText[upperCaseStat].In.text}}</p>
                 </v-row>
          <div v-for="part in textparts" :key="part">
            <v-row justify="center">
                 <h1>{{statText[upperCaseStat][part].title}}</h1>
                 </v-row>
                    <v-row justify="center" >
             <p>{{statText[upperCaseStat][part].text}}</p>
               
                    </v-row>
            
             </div>
             
           
            <v-row justify="center">
             
            </v-row>

<!--
              <v-row justify="center">
                
                   <v-col cols="6" justify="center">
                         <div v-if="!bplantable">
                       <p  >No free spaces in this module!</p>
                       </div>
                   </v-col>
                    <v-col  cols="6"  justify="center">
                        <div v-if="!bharvestable">
                         <p  >Nothing to harvest in this module!</p>
                        </div>
                   </v-col>
              </v-row>
-->
           
             
            
        </div>
    </FarmTransition>
</template>

<script>
    import FarmTransition from "../../main/FarmTransitionStatInfo"
    import {mapState} from "vuex"
    import StatText from "@/assets/stat_info/StatText.json"
 

    export default {
        name: "StatInfo",
        components: {
            FarmTransition,
        },
        props: {
            InfoType: String,
            
        },
        data() {
            return {
                yPositions: [260, 0, -300, -600],
                statText: StatText,
                textparts: ["Ts", "Nu", "Ku"],
              
            }
        },
         computed: {
            ...mapState(["farm_info"]),
            upperCaseStat: function () {
                return this.InfoType.replace(/\b\w/g, l => l.toUpperCase())
            },
           
        
        },
        methods: {

             getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic.replace(/\b\w/g, l => l.toUpperCase())+".png")
            },
            setYpos: function(){
                this.$store.dispatch('change_ypos_statInfo', 260);
            },
            upYpos: function(){
                this.$store.dispatch('change_ypos_statInfo', -600);
               
                /*this.yPos -= 200
                if(this.yPos < -600){
                    this.yPos = -650
                }*/
            },
           
        },
      

    }
</script>

<style scoped>
    h2, h1{
        font-weight: 600;
        font-size: 36px;
        line-height: 44px;
        padding-top: 10px;
        color: var(--v-primary-base);
    }

    p{
        text-align: center;
        width: 410px;
        color: var(--v-primary-base);
        padding-top: 10px;
        margin: 0 !important;
    }

    #button{
        font-style: normal;
        font-weight: bold;
        font-size: 24px;
        line-height: 29px;
        border-radius: 47px;
    }

    .info-box{
        background: #789659;
        border-radius: 10px;
        box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
        margin: 40px 15px 0 15px;
    }
  
</style>