<template>
    <FarmTransition :y-positions="yPositions" :yPos="yPos">
        <div>
            <v-row justify="center"  aling-self="center">
                <v-col cols="1">
                  <v-btn dark fab small color="white" @click="upYpos()">
            <v-icon color="primary">mdi-arrow-up</v-icon>
          </v-btn></v-col> 
              <v-col cols="9" justify="center">  <h2>Module {{InfoType}}: {{upperCasePlant}}</h2></v-col>
           <v-col cols="1">      <v-btn dark fab small color="white" @click="setYpos()">
            <v-icon color="primary">mdi-close</v-icon>
          </v-btn></v-col>
            </v-row>
            <v-row justify="center">
          
            </v-row>
            <v-row justify="center">
                <img :src="getImgUrl(module_plants.type)" alt="" width="120px" height="auto">
        
            </v-row>


        <v-row justify="space-around">   
            <v-col cols="4">
                <v-btn id="button" :disabled="!bplantable" class="text-capitalize" rounded color="accent" height="60" width="150" @click="PlantHere()">
                    Plant
                    <v-icon right dark>mdi-arrow-right</v-icon>
                 </v-btn>
            </v-col>

            <v-vcol cols="4">
                 <v-btn id="button" :disabled="!bharvestable" class="text-capitalize" rounded color="accent" height="60" width="150" @click="HarvestHere()">
                    Harvest
                    <v-icon right dark>mdi-arrow-right</v-icon>
                </v-btn>
            </v-vcol>
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
             <div v-for="part in textparts" :key="part">
            <v-row justify="center">
                 <h1>{{plantText[upperCasePlant][part].title}}</h1>
                 </v-row>
                    <v-row justify="center" >
             <p>{{plantText[upperCasePlant][part].text}}</p>
               
                    </v-row>
            
             </div>
             
            
        </div>
    </FarmTransition>
</template>

<script>
    import FarmTransition from "../main/FarmTransition";
    import {mapState, mapGetters} from "vuex"
    import PlantText from "@/assets/plant_info/PlantText.json"
    import axios from "axios"
 

    export default {
        name: "StatExplain",
        components: {
            FarmTransition,
        },
        props: {
            InfoType: Number,
            yPos: Number,
        },
        data() {
            return {
                yPositions: [260, 0, -300, -600],
                plantText: PlantText,
                textparts: ["In", "Ts", "Nu", "Ku"],
                plantable: null,
                plantamodules: null
            }
        },
         computed: {
            ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.InfoType]
            },
            upperCasePlant: function () {
                return this.module_plants.type.replace(/\b\w/g, l => l.toUpperCase())
            },
            bharvestable: function () {
                if(this.module_plants.pos.find(o => o.harvestable === true)){
                return true}
                else{
                    return false
                }
                //return this.harvestable.find(o => o.plant_type === 'Mint')
    
            },
           

            bplantable: function () {
                if(this.module_plants.pos.find(o => o.age === 0) && this.module_plants.pos.find(o => o.harvestable === null) && this.plantamodules.find(o => o.available_plants === this.InfoType)){
                return true}
                else{
                    return false
                }
            },
        
        },
        methods: {
            findOldestPlant(){
                var highage = this.module_plants.pos[5].age
                var pos = 5
                for (var i = 5; i >= 0; i--){
                    
                    if(this.module_plants.pos[i].age > highage){
                        pos = i
                    }
                }
                return pos
            },
             getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic.replace(/\b\w/g, l => l.toUpperCase())+".png")
            },
            setYpos: function(){
                this.yPos = 260
            },
            upYpos: function(){
                this.yPos = -600
                /*this.yPos -= 200
                if(this.yPos < -600){
                    this.yPos = -650
                }*/
            },
            BlinkModule: function () {
                axios.post("http://127.0.0.1:3000/dashboard/blink",
                    {module: this.InfoType},
                    "content-type: application/json")
                    .then()
                    .catch(error => {
                        console.log(error);
                    });
            },
               getPlantablePlants: function () {
                axios.get("http://127.0.0.1:3000/dashboard/plantable-plants")
                    .then(result => {
                        this.plantable = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
             getPlantableModules: function () {
                axios.get("http://127.0.0.1:3000/dashboard/plantable-modules")
                    .then(result => {
                        this.plantamodules = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
             ...mapGetters(["get_module"]),
             PlantHere: function(){
                this.$store.dispatch('insertFarming', {m: this.InfoType,p: 0,t:  this.module_plants.type.replace(/\b\w/g, l => l.toUpperCase())});
                this.BlinkModule()
                this.$router.push('/plant');
                
             },
             HarvestHere: function(){
                this.$store.dispatch('insertFarming', {m: this.InfoType,p: this.findOldestPlant(),t:  this.module_plants.type.replace(/\b\w/g, l => l.toUpperCase())});
                this.BlinkModule()
                this.$router.push('/harvest');
             }

        },
        created() {
           this.getPlantablePlants()
           this.getPlantableModules()
          
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