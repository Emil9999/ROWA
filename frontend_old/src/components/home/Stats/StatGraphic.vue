<template>
    <div style="padding: 0 30px">
        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'temp')">
                <StatBox heading="Temperature" type="temperature" :value="temperature"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'light_intensity')">
                <StatBox heading="Box Temperature" type="e-temperature" :value="light_intensity"></StatBox>
            </v-col>
        </v-row>

        <v-row>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'water_temp')">
                <StatBox heading="Water Temperature" type="temperature" :value="water_temp"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'humidity')">
                <StatBox heading="Humidity" type="humidity" :value="humidity"></StatBox>
            </v-col>
        </v-row>


        <v-row>
               <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'water_level')">
                <StatBox heading="Water Level" type="waterLevel" :value="water_level"></StatBox>
            </v-col>
            <v-col style="padding: 15px" v-on:click="$emit('infoOn', 'water_ph')">
                <StatBox heading="pH" type="ph" :value="water_ph"></StatBox>
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
                temperature: null,
                light_intensity: null,
                interval: null,
                humidity: null,
                water_level: null,
                water_temp: null,
                water_ph: null,
                max_height: 23,
                min_height: 12
            }
        },
        methods: {
            getSensorData: function () {
                axios.get("http://localhost:3000/dashboard/sensor-data").then(response => {
                    this.temperature = response.data.temperature
                    this.light_intensity = response.data.light_intensity
                    this.humidity = response.data.humidity
                    this.water_level = Math.round(((this.max_height - response.data.water_level)/(this.max_height -  this.min_height))*100)
                    this.water_temp = response.data.water_temp
                    this.water_ph = response.data.water_ph
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