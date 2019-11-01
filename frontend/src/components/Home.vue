<template>
  <div class="container">
    <div class="row justify-content-xl-center">
      <img src="../assets/logo_rowa.png"/>
      <hr/>
    </div>
    <div class="row">
      <h4>Total plants:</h4>
<!--      <h4 v-for="plant for harvestable_plants"> {{ harvestable_plants }}</h4>-->
    </div>
    <div class="row" v-for="(plant, index) in all_plants" v-bind:key="index">
     Module {{plant.position}} {{plant.plant_type}}: {{plant.available_plants}}
    </div>
    <div class="row">
      <h4>Plants to harvest:</h4>
<!--      <h4 v-for="plant for harvestable_plants"> {{ harvestable_plants }}</h4>-->
    </div>
    <div class="row" v-for="(plant, index) in harvestable_plants" :key="`plant-${index}`">
      {{plant.plant_type}}: {{plant.available_plants}}
    </div>
    <div class="row">
      <h1>{{ msg }}</h1>
    </div>
    <div class="row">
      <router-link to="/plant">
        <button type="button" class="btn btn-primary">Plant</button>
      </router-link>
      <router-link to="/harvest">
        <button type="button" class="btn btn-primary">Harvest</button>
      </router-link>
    </div>
    <div class="row">
      <h2>Sensor Data</h2>
    </div>
    <div class="row">
      Temperature: {{sensor_data.temperature}} °C, Light Intensity: {{sensor_data.light_intensity}} Lux, Updated: {{sensor_data_updated}}
    </div>
  </div>
</template>

<script>
    import axios from "axios"

    export default {
        name: "Home",
        data() {
            return {
                msg: "Farm Overview",
                harvestable_plants: null,
                all_plants:null,
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
            getHarvestablePlants: function() {
               axios.get("http://127.0.0.1:3000/dashboard/harvestable")
                .then(result => {
                    this.harvestable_plants = result.data
                })
                .catch(error => {
                    console.log(error)
                })
            },
            getAllPlants: function(){
               axios.get("http://127.0.0.1:3000/dashboard/all")
                .then(result => {
                    this.all_plants = result.data
                })
                .catch(error => {
                    console.log(error)
                })

            }
        },
        created() {
          //Get request to get all plants in farm + harvestable plants
          this.getHarvestablePlants()
          this.getAllPlants()

          // Get sensor data and request them every 10 Seconds
          this.getSensorData()
          this.interval = setInterval(this.getSensorData, 10000);
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
