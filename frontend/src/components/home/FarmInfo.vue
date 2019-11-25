<template>
    <FarmTransition :y-positions="yPositions">
        <div>
            <v-row justify="center">
                <h2>Farming</h2>
            </v-row>
            <v-row justify="center">
                <p>Plant and Harvest your very own plant from this farm for your lunch, or to take home.</p>
            </v-row>
            <v-row justify="center" style="margin-top: 40px">
                 <router-link :to="{name:'Farming'}">
                <v-btn id="button" class="text-capitalize" rounded color="accent" height="75" width="360" @click.stop="hello">
                    Start Farming Now
                    <v-icon right dark>mdi-arrow-right</v-icon>
                </v-btn>
                 </router-link>
            </v-row>
            <v-row style="padding: 0 80px">
                <v-col class="info-box">
                    <InfoBoxPlants heading="Harvestable" :plants="harvestable_plants"></InfoBoxPlants>
                </v-col>
                <v-col class="info-box">
                    <InfoBoxPlants heading="Plantable" :plants="harvestable_plants"></InfoBoxPlants>
                </v-col>
            </v-row>
        </div>
    </FarmTransition>
</template>

<script>
    import FarmTransition from "../main/FarmTransition";
    import InfoBoxPlants from "./InfoBoxPlants";

    export default {
        name: "FarmInfo",
        components: {
            FarmTransition,
            InfoBoxPlants
        },
        props:{
            harvestable_plants: Array
        },
        data() {
            return {
                yPositions: [0, -160]
            }
        },
        methods: {
            hello: function () {
                console.log("test")
            }
        }
    }
</script>

<style scoped>
    h2{
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