<template>
    <FarmTransition :y-positions="yPositions" :yPos="yPos">
        <div>
            <v-row justify="center">
                <h2>Module: {{InfoType}}</h2>
               
            </v-row>
            <v-row justify="center">
            <h1>{{module_plants.type.replace(/\b\w/g, l => l.toUpperCase())}}</h1>
            </v-row>
            <v-row justify="center">
                <img :src="getImgUrl(module_plants.type)" alt="" width="120px" height="auto">
            </v-row>
            <v-row justify="center" style="margin-top: 40px">
             <h2></h2>
                <v-btn id="button" class="text-capitalize" rounded color="accent" height="75" width="360" @click.stop="hello" :to="{name:'Farming'}">
                    Start Farming Now
                    <v-icon right dark>mdi-arrow-right</v-icon>
                </v-btn>
      
            </v-row>
            <v-row style="padding: 0 80px">
                <v-col class="info-box">
                    <InfoBoxPlants heading="Harvestable" :plants="harvestable"></InfoBoxPlants>
                </v-col>
                <v-col class="info-box">
                    <InfoBoxPlants heading="Plantable" :plants="plantable"></InfoBoxPlants>
                </v-col>
            </v-row>
        </div>
    </FarmTransition>
</template>

<script>
    import FarmTransition from "../main/FarmTransition";
    import InfoBoxPlants from "./InfoBoxPlants";
    import axios from "axios"
    import {mapState, mapGetters} from "vuex"
    import PlantText from "@/assets/plant_info/PlantText.json"

    export default {
        name: "StatExplain",
        components: {
            FarmTransition,
            InfoBoxPlants
        },
        props: {
            InfoType: Number,
            yPos: Number,
        },
        data() {
            return {
                yPositions: [260, 0, -260],
                harvestable: null,
                plantable: null,
                plantText: PlantText
            }
        },
         computed: {
            ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.InfoType]
            },
        
        },
        methods: {
             getImgUrl(pic) {
                return require('@/assets/harvesting/plants/'+pic.replace(/\b\w/g, l => l.toUpperCase())+".png")
            },
            getTextUrl(part){
                return plantText.[this.module_plants.type].In.text
            },
             ...mapGetters(["get_module"]),
            getHarvestablePlants: function () {
                axios.get("http://127.0.0.1:3000/dashboard/harvestable-plants")
                    .then(result => {
                        this.harvestable = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
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

        },
        created() {
            this.getHarvestablePlants()
            this.getPlantablePlants()
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