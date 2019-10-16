<template>
  <div class="container">
    <div class="row justify-content-xl-center">
      <h1>Plant</h1>
      <img src="../assets/logo_rowa.png" />
    </div>
    <div class="row">
      <h2>Choose a plant type</h2>
    </div>
    <div class="row">
      <!-- Get available plant types from db and generate list based on that-->
      <ul id="example-1">
        <li v-for="(plant, index) in plantTypes" v-bind:key="index">
          <b-button v-on:click="postreq">{{plant}}</b-button>
        </li>
      </ul>
      
    </div>

    <div class="row">
      <h2>Please plant in position {{position}}</h2>
    </div>
    <div class="row">
     <router-link to="/plant/howto">
        <button type="button" class="btn btn-primary" v-on:click="sendPosition">Plant Now</button>
      </router-link>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import { eventBus } from '../main';
export default {
  data() {
    return {
      plantTypes:[],
      position: null
    };
  },
  methods: {
    postreq: function() {
      axios
        .post(
          "http://127.0.0.1:3000/plant/plant",
          { plantType: "Lettuce" },
          "content-type: application/json"
        )
        .then(result => {
          this.position = result.data;
        })
        .catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });
    },
    sendPosition: function(){
      eventBus.$emit('sendPosition', this.position)
    }
  },
  created(){
    axios.get("http://127.0.0.1:3000/plant/available")
    .then(result => {
      console.log(JSON.parse(result.data))
      this.plantTypes =JSON.parse(result.data)
    })
    .catch(error => {
      console.log(error)
    })
  }
};
</script>