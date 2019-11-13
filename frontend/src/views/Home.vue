<template>
    <div id="Home">
        <v-container>
            <HomeTopRow></HomeTopRow>
        </v-container>
        <FarmInfo></FarmInfo>
    </div>
</template>

<script>
    import axios from "axios"
    import HomeTopRow from "@/components/main/HomeTopRow"
    import FarmInfo from "@/components/home/FarmInfo";

    export default {
        name: "Home",
        components: {
            HomeTopRow,
            FarmInfo
        },
        data() {
            return {
                msg: "Farm Overview",
                harvestable_plants: null,
                all_plants: null,
                sensor_data: {
                    datetime: null,
                    temperature: null,
                    light_intensity: null
                },
                sensor_data_updated: null,
                interval: null
            };
        },
        methods: {
            getSensorData: function () {
                this.sensor_data_updated = new Date().toISOString()
                axios.get("http://127.0.0.1:3000/dashboard/sensor-data")
                    .then(result => {
                        this.sensor_data = result.data[0]
                        console.log(this.sensor_data)
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
            getHarvestablePlants: function () {
                axios.get("http://127.0.0.1:3000/dashboard/harvestable")
                    .then(result => {
                        this.harvestable_plants = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
            },
            getAllPlants: function () {
                axios.get("http://127.0.0.1:3000/dashboard/all")
                    .then(result => {
                        this.all_plants = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })

            },
            getIntervalData: function () {
                this.interval = setInterval(this.getSensorData, 5000);
            }
        },
        created() {
            //Get request to get all plants in farm + harvestable plants
            this.getHarvestablePlants()
            this.getAllPlants()

            // Get sensor data and request them every 10 Seconds
            this.getSensorData()
            // this.getIntervalData()
        },
        beforeDestroy() {
            clearInterval(this.interval)
        }
    };
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    h1,
    h2 {
        font-weight: normal;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }
</style>
