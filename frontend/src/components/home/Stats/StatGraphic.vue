<template>
    <div style="padding: 0 30px">
        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'externalTemp')">
                <StatBox heading="Temperature" type="temperature" :value="externalTemp"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'boxTemp')">
                <StatBox heading="Box Temperature" type="e-temperature" :value="boxTemp"></StatBox>
            </v-col>
        </v-row>

        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'externalHumidity')">
                <StatBox heading="Humidity" type="humidity" :value="externalHumidity"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'boxHumidity')">
                <StatBox heading="Box Humidity" type="humidity" :value="boxHumidity"></StatBox>
            </v-col>
        </v-row>


        <v-row>
               <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'water_level')">
                <StatBox heading="Water Level" type="waterLevel" :value="water_level"></StatBox>
            </v-col>
        </v-row>

        <v-row justify="center">
            <p></p>
        </v-row>
    </div>
</template>

<script>
    import StatBox from "./StatBox";
    import axios from "axios"

    export default {
        name: "StatGraphic",
        components: {
            StatBox
        },
        data() {
            return {
                externalTemp: null,
                interval: null,
                externalHumidity: null,
                water_level: null,
                boxHumidity: null,
                boxTemp: null,
            }
        },
        methods: {
            getSensorData: function () {
                axios.get("http://localhost:3000/dashboard/sensor-data").then(response => {
                    this.externalTemp = response.data.externalTemp
                    this.externalHumidity = response.data.externalHumidity
                    this.water_level = response.data.water_level
                    this.boxHumidity = response.data.boxHumidity
                    this.boxTemp = response.data.boxTemp
                    console.log(response.data)
                })
            }
        },
        mounted() {
            this.getSensorData()
            this.interval = setInterval(this.getSensorData, 30000)
        },
        destroyed() {
            clearInterval(this.interval)
        }

    }
</script>

<style scoped>


</style>