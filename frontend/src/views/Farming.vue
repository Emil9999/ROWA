<template>
  <div id="Farming">
    <v-container align="center">
      <FarmingTopRow v-bind:step="step" v-bind:headtext="'Choose what you would like to do!'"></FarmingTopRow>

      <v-row justify="center">
        <h1>Harvest a plant</h1>
      </v-row>

      <v-row justify="center">
        <p>Go on a journey to become a famer.Harvest a plant now, for your daily dose of fresh greens.</p>
      </v-row>
      <v-row justify="center" class="info-box" margin="auto" padding="auto">
        <InfoBoxPlants heading="Harvestable" :plants="harvestable_plants"></InfoBoxPlants>
      </v-row>

      <v-row justify="center" style="margin-top: 40px">
        <router-link :to="{name:'Harvest'}">
          <v-btn id="button" rounded color="accent" height="75" width="360">
            Start Harvesting
            <v-icon right dark>mdi-arrow-right</v-icon>
          </v-btn>
        </router-link>
      </v-row>

      <v-row justify="center">
        <h1>Plant a new Plant</h1>
      </v-row>

      <v-row justify="center">
        <p>Go on a journey to become a famer.Plant a new seed now and enjoy fresh greens in the future.</p>
      </v-row>
      <v-row justify="center" class="info-box" margin="auto" padding="auto">
        <InfoBoxPlants heading="Plantable" :plants="plantable_plants"></InfoBoxPlants>
      </v-row>

      <v-row justify="center" style="margin-top: 40px">
        <router-link :to="{name:'Plant'}">
        <v-btn id="button" rounded color="accent" height="75" width="360">
          Start Planting
          <v-icon right dark>mdi-arrow-right</v-icon>
        </v-btn>
        </router-link>
      </v-row>
    </v-container>
  </div>
</template>

<script>
import axios from "axios";
import FarmingTopRow from "../components/farming/FarmingTopRow";
import InfoBoxPlants from "../components/home/InfoBoxPlants";
export default {
  name: "Farming",
  components: {
    FarmingTopRow,
    InfoBoxPlants
  },
  data() {
    return {
      step: 0,
      harvestable_plants: null,
      plantable_plants: null
    };
  },
  methods: {
    getHarvestablePlants: function() {
      axios
        .get("http://127.0.0.1:3000/dashboard/harvestable")
        .then(result => {
          this.harvestable_plants = result.data;
          this.plantable_plants = result.data;
        })
        .catch(error => {
          console.log(error);
        });
    }
  },
  created() {
    //Get request to get all plants in farm + harvestable plants
    this.getHarvestablePlants();
  }
};
</script>

<style scoped>
#button {
  font-weight: bold;
  margin: 40px;
  font-family: Montserrat;
  font-size: 24px;
}
h1 {
  color: #789659;
}

p {
  text-align: center;
  width: 410px;
  color: #828282;
  padding-top: 10px;
  margin: 0 !important;
}
.info-box {
  background: #789659;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 100px 0 100px;
}
</style>