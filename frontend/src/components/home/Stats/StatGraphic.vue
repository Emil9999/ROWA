<template>
    <div style="padding: 0 30px">
        <v-row>
            <v-col style="padding: 15px">
                <StatBox heading="Temperature" type="temperature" :value="temperature"></StatBox>
            </v-col>
            <v-col style="padding: 15px">
                <StatBox heading="Light Intensity" type="light" :value="light_intensity"></StatBox>
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
                interval: null
            }
        },
        methods: {
            getSensorData: function () {
                axios.get("http://localhost:3000/dashboard/sensor-data").then(response => {
                    this.temperature = response.data.temperature
                    this.light_intensity = response.data.light_intensity
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