<template>
  <div class="container">
    <div class="row justify-content-xl-center">
      <img src="../assets/logo_rowa.png"/>
      <hr/>
    </div>
    <div class="row">
      <h4>Plants to harvest:</h4>
<!--      <h4 v-for="plant for harvestable_plants"> {{ harvestable_plants }}</h4>-->
    </div>
    <div class="row" v-for="plant in harvestable_plants">
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
  </div>
</template>

<script>
    import axios from "axios"

    export default {
        name: "Home",
        data() {
            return {
                msg: "Farm Overview",
                harvestable_plants: null
            };
        },
        created() {
            axios.get("http://127.0.0.1:3000/dashboard/main")
                .then(result => {
                    console.log(JSON.parse(result.data))
                    this.harvestable_plants = JSON.parse(result.data)
                })
                .catch(error => {
                    console.log(error)
                })
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
