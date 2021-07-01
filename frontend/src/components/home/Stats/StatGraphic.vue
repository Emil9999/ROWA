<template>
    <div style="padding: 0 30px">
        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'temp')">
                <StatBox heading="Temperature" type="temperature" :value="externalTemp"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'light_intensity')">
                <StatBox heading="Box Temperature" type="e-temperature" :value="boxTemp"></StatBox>
            </v-col>
        </v-row>

        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'water_temp')">
                <StatBox heading="Humidity" type="humidity" :value="externalHumidity"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'humidity')">
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
                max_height: 23,
                min_height: 12
            }
        },
        methods: {
            getSensorData: function () {
                axios.get("http://localhost:3000/dashboard/sensor-data").then(response => {
                    this.externalTemp = response.data.externalTemp
                    this.externalHumidity = response.data.externalHumidity
                    this.water_level = Math.round(((this.max_height - response.data.water_level)/(this.max_height -  this.min_height))*100)
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