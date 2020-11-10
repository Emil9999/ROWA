<template>
    <FarmTransition :y-positions="yPositions" :yPos="yPos">
        <div>
            <v-row justify="center">
                <h2>Module: {{InfoType}}</h2>
               
            </v-row>
            <v-row justify="center">
            <h1>{{upperCasePlant}}</h1>
            </v-row>
            <v-row justify="center">
                <img :src="getImgUrl(module_plants.type)" alt="" width="120px" height="auto">
                <p>Harvest: {{bharvestable}}</p>
                <p>Plant: {{bplantable}}</p>
            </v-row>
            <v-row style="padding: 0 80px">
                 <v-btn id="button" class="text-capitalize" rounded color="accent" height="75" width="360" @click="PlantHere()">
                    Start Farming Now
                    <v-icon right dark>mdi-arrow-right</v-icon>
                </v-btn>
                
            </v-row>
             <div v-for="part in textparts" :key="part">
            <v-row justify="center" style="margin-top: 40px">
                 <h1>{{plantText[upperCasePlant][part].title}}</h1>
                 </v-row>
                    <v-row justify="center" style="margin-top: 40px">
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
    import {eventBus} from "@/main.js";

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
                yPositions: [260, 0, -260],
                plantText: PlantText,
                textparts: ["In", "Ts", "Nu", "Ku"]
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
                if(this.module_plants.pos.find(o => o.age === 0) && this.module_plants.pos.find(o => o.harvestable === null)){
                return true}
                else{
                    return false
                }
            }
        
        },
        methods: {
             getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic.replace(/\b\w/g, l => l.toUpperCase())+".png")
            },
             ...mapGetters(["get_module"]),
             PlantHere: function(){
                
                 this.$router.push('/plant')
                eventBus.$emit('planthere');
             }

        },
        created() {
           
        }
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