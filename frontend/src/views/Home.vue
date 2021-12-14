<template>
  <div id="Home">
    <v-container>
      <HomeTopRow />
      <div v-if="farm_active==0">
        <v-row style="margin: 30px 0 0 50px">
          <v-col>
            <v-chip color="transparent">
              <v-avatar color="primary" size="16px" style="margin-right: 7px"></v-avatar>
              <span>
                Ready to
                <br />harvest
              </span>
            </v-chip>
            <v-chip color="transparent">
              <v-avatar color="accent" size="16px" style="margin-right: 7px"></v-avatar>
              <span>
                Ready to
                <br />plant
              </span>
            </v-chip>
          </v-col>
        </v-row>
        <v-row justify="center" style="padding-top: 40px">
          <CatTree v-on:moduleClicked="zoomToModule" />
        </v-row>
      </div>
      <div v-if="farm_active==1">
        <div class="harvest-box" justify="center">
          <v-row justify="center" style="margin: 0px 0px 20px 0px;">
            <h1 style="color:primary" >Welcome to the <br> Susteyn Office Farm</h1>
          </v-row>
          <v-row justify="center">
            <p
              style="text-allign:block"
            >The Office Farm is your compagnion <br> for healthier lunch and a greener workspace.</p>
          </v-row>
          <v-row justify="center" style="margin: 20px 0px 20px 0px;">
            <h2 style="">How does it work?</h2>
          </v-row>
          <v-row justify="center">
            <ul>
              <li>Everyone (you) can interact with the Farm.</li>
              <li>Each week there is something for you to harvest and new plants to plant.</li> 
              <li>Follow the instructions on the screen to find out what you can plant and what you can harvest.</li>
            </ul>
          </v-row>
        </div>
         <v-row justify="center" style="margin-top: 20px">
           <h1>What do you want to do?</h1>
         </v-row>
        <v-row justify="center" style="margin-top: 70px">
          <v-btn
            id="button"
            :disabled="HarvestDisable"
            :to="{name:'Harvest'}"
            rounded
            color="primary"
            height="75"
            width="360"
          >
            Harvest
            <v-icon right dark>mdi-arrow-right</v-icon>
          </v-btn>
        </v-row>
        <v-row justify="center" style="margin-top: 50px">
          <v-btn
            id="button"
            :disabled="PlantDisable"
            :to="{name:'Plant'}"
            rounded
            color="primary"
            height="75"
            width="360"
          >
            Plant
            <v-icon right dark>mdi-arrow-right</v-icon>
          </v-btn>
        </v-row>
      </div>
      <div v-if="farm_active==2">
        <StatGraphic v-on:infoOn="onInfoOn"></StatGraphic>
      </div>
    </v-container>
    <div v-if="farm_active!=1">
      <FarmInfo />
    </div>
    <div>
      <StatInfo v-bind:InfoType="this.info_type_stat" />
    </div>
    <div>
      <PlantInfo v-bind:InfoType="this.info_type" />
    </div>
  </div>
</template>

<script>
import axios from "axios";
import HomeTopRow from "@/components/main/HomeTopRow";
import FarmInfo from "@/components/home/FarmInfo";
import { mapState } from "vuex";
import CatTree from "../components/home/Farm/CatTree";
import StatGraphic from "../components/home/Stats/StatGraphic";
import PlantInfo from "@/components/home/PlantInfo";
import StatInfo from "@/components/home/Stats/StatInfo";

export default {
  name: "Home",
  components: {
    HomeTopRow,
    FarmInfo,
    CatTree,
    StatGraphic,
    PlantInfo,
    StatInfo
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
      interval: null,
      moduleNumber: null,
      reverse: false,
      info_type: 1,
      info_type_stat: "ExternalTemp",
      harvestable_plant: null,
      plantable_plants: null
    };
  },
  methods: {
    onInfoOn: function(type) {
      this.info_type_stat = type;
      this.$store.dispatch(
        "change_ypos_statInfo",
        0 + Math.floor(Math.random() * 10) / 1000
      );
    },

    getSensorData: function() {
      this.sensor_data_updated = new Date().toISOString();
      axios
        .get("http://127.0.0.1:3000/dashboard/sensor-data")
        .then(result => {
          this.sensor_data = result.data[0];
        })
        .catch(error => {
          console.log(error);
        });
    },
    getIntervalData: function() {
      this.interval = setInterval(this.getSensorData, 5000);
    },
    getHarvestablePlants: function() {
      axios
        .get("http://127.0.0.1:3000/dashboard/harvestable-plants")
        .then(result => {
          this.harvestable_plant = result.data;
        })
        .catch(error => {
          console.log(error);
        });
    },
    getPlantablePlants: function() {
      axios
        .get("http://127.0.0.1:3000/dashboard/plantable-plants")
        .then(result => {
          this.plantable_plants = result.data;
          console.log(this.plantable_plants);
        })
        .catch(error => {
          console.log(error);
        });
    },
    zoomToModule: function(moduleNumber) {
      this.moduleNumber = moduleNumber;
      this.reverse = moduleNumber % 2 === 0;
      this.info_type = moduleNumber;
      this.$store.dispatch(
        "change_ypos_plantInfo",
        -50 + (Math.floor(Math.random() * 10) + moduleNumber) / 1000
      );
    }
  },
  created() {
    // Get sensor data and request them every 10 Seconds
    this.getSensorData();
    this.getHarvestablePlants();
    this.getPlantablePlants();
    // this.getIntervalData()
  },
  beforeDestroy() {
    clearInterval(this.interval);
  },
  computed: {
    ...mapState(["farm_active"]),
    PlantDisable: function() {
      if (this.plantable_plants == null) {
        return true;
      } else {
        return false;
      }
    },
    HarvestDisable: function() {
      if (this.harvestable_plant == null) {
        return true;
      } else {
        return false;
      }
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1,
h2 {
  font-weight: normal;
  font-family: Montserrat;
  font-style: normal;
  font-weight: 600;
  text-align: center;

  color: #A0918D;
}
h1 {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 600;
  text-align: center;

  color: var(--v-primary-base);
}
ul {
  
}
p{
  color: #A0918D;
  text-align: center;
}

li {
  color: #A0918D;
}

a {
  color: #42b983;
}

span {
  color: var(--v-primary-base);
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
}

.no-hover:hover {
  background-color: transparent;
  text-decoration: none;
}
.harvest-box {
  background: #ffffff;
  border-radius: 20px;
  margin: 60px 20px 120px 20px;
  padding: 50px;
}
#logo-text {
  font-family: Montserrat, sans-serif;
  font-style: normal;
  font-weight: 600;
  font-size: 32px;
  color: var(--v-primary-base);
}
#button {
  font-style: normal;
  font-weight: bold;
  font-size: 24px;
  line-height: 29px;
  border-radius: 47px;
}
</style>
