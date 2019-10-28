<template>
  <div>
    <h1>Harvest</h1>
    <ul id="example-1">
        <li v-for="(plant, index) in plantTypes" v-bind:key="index">
          <b-button v-on:click="postreq(plant)">{{plant}}</b-button>
        </li>
      </ul>
      <h3>Please harvest from module {{modulePosition}} in position {{plantPosition}}</h3>
       <b-button :disabled="isDisabled" v-on:click="harvestDone">Harvest now</b-button>
  </div>
</template>
<script>
import axios from "axios";
    export default {
      data() {
        return {
          isDisabled: true,
          plantTypes:[],
          plantPosition:null,
          modulePosition:null
        }
      },
      created(){
        axios.get("http://127.0.0.1:3000/harvest/available")
        .then(result => {
          console.log(result.data)
         this.plantTypes =result.data
        })
        .catch(error => {
          console.log(error)
        })
      },
      methods: {
        harvestDone: function(){
          axios.post(
          "http://127.0.0.1:3000/harvest/harvestdone",
          { plantPosition: this.plantPosition, modulePosition:this.modulePosition },
          "content-type: application/json"
        )
        .then(result => {
          console.log(result.data)
          this.$router.push('/')
          
        })
        .catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });
        },
        postreq: function(selectedPlant) {
          this.isDisabled = false
      axios.post(
          "http://127.0.0.1:3000/harvest/harvest",
          { plantType: selectedPlant },
          "content-type: application/json"
        )
        .then(result => {
          console.log(result.data)
          this.plantPosition = result.data[0];
          this.modulePosition = result.data[1];
        })
        .catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });
      }
      }

    }
</script>
