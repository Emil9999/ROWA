<template>
  <div id="BatchFarming">
    <AdminTopRow v-bind:headtext="'Batch Farming'" v-bind:prevPage="'AdminMenu'"></AdminTopRow>
    <BatchButtonPill></BatchButtonPill>
     <div v-if="batch_active">
                 <div  v-for="item in harvestableModules" :key="item.module_position">
       
   <v-row class="info-box" justify="center"> 
     
       <v-col align-self="center" align="center"> <h3>{{item.plant_type}} Module: {{item.module_position}} Position: {{item.plant_position}}</h3>  </v-col>
        <v-col> <img :src="getImgUrl(item.plant_type)" alt="" width="120px" height="auto"> </v-col>
        <v-checkbox v-model="selectedHarvestable" label="Select" :value="item"></v-checkbox>
     
   </v-row>
   
    </div>
                <v-btn id="button" rounded color="accent" height="50" width="360" @click="sendHarvesting()">Save 
                  <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
                </v-btn>
            </div>
            <div v-else>

              <div  v-for="item in plantableModules" :key="item.planted_module">
       
   <v-row class="info-box" justify="center"> 
     
       <v-col align-self="center" align="center"> <h3>{{item.plant_type}} Module: {{item.planted_module}} </h3>  </v-col>
        <v-col> <img :src="getImgUrl(item.plant_type)" alt="" width="120px" height="auto"> </v-col>
        <v-checkbox v-model="selectedPlantable" label="Select" :value="item.planted_module"></v-checkbox>
     
   </v-row>
   
    </div>
                <v-btn id="button" rounded color="accent" height="50" width="360" @click="sendPlanting()">Save 
                  <v-icon right dark>mdi-checkbox-marked-circle</v-icon>
                </v-btn>
            </div>
  </div>
</template>


<script>
import AdminTopRow from "@/components/admin/AdminTopRow.vue";
import BatchButtonPill from "@/components/admin/BatchButtonPill.vue";
import axios from "axios";
import {mapState} from "vuex"
export default {
  name: "BatchFarming",
  components: {
    AdminTopRow,
    BatchButtonPill
  },
  data() {
    return {
      plantableModules: null,
      harvestableModules: null,
      selectedPlantable: [],
      selectedHarvestable: []
    };
  },
  methods: {
    getImgUrl(pic) {
      return require("@/assets/harvesting/plants/" + pic + ".png");
    },
    sendPlanting: function() {
      axios
        .post(
          "http://127.0.0.1:3000/plant/plant-all",
          {planted_modules: this.makeString(this.selectedPlantable) },
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
    },
    sendHarvesting: function() {
      axios
        .post(
          "http://127.0.0.1:3000/harvest/harvest-all",
          this.selectedHarvestable,
          "content-type: application/json"
        )
        .then()
        .catch(error => {
          console.log(error);
        });
    },
    getPlantable: function() {
      axios
        .get("http://127.0.0.1:3000/plant/get-all")
        .then(result => {
          this.plantableModules = result.data;
        })
        .catch(error => {
          console.log(error);
        });
    },
    getHarvestable: function() {
      axios
        .get("http://127.0.0.1:3000/harvest/get-all")
        .then(result => {
          this.harvestableModules = result.data;
         
        })
        .catch(error => {
          console.log(error);
        });
    },
    makeString: function(NumberArray){
      var NumberString = ","
      for (var number of NumberArray) {
          NumberString = NumberString + number.toString() + ","
    }
      return NumberString
    }
  },
  computed: {
          ...mapState(["batch_active"])
  },
  created() {
    this.getPlantable();
    this.getHarvestable();
  }
};
</script>

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

tr {
  font-family: Montserrat;
  font-style: normal;
  font-weight: 500;
  font-size: 32px;
  line-height: 29px;
  color: var(--v-primary-base);
}
span {
  color: var(--v-primary-base);
  font-style: normal;
  font-weight: normal;
  font-size: 14px;
}

.carussel {
  color: var(--v-primary-base);
}

.info-box {
  background: #ffffff;
  border-radius: 10px;
  box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
  margin: 40px 50px 0 50px;
}
</style>