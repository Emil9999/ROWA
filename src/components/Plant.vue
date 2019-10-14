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
  </div>
</template>
<script>
import axios from "axios";

export default {
  data() {
    return {
      plantTypes:["Lettuce","Basil"], //This is supposed to come from the server
      position: null
    };
  },
  methods: {
    postreq: function() {
      axios
        .post(
          "http://127.0.0.1:3000/plant/",
          { plantType: "Lettuce" },
          "content-type: application/json"
        )
        .then(result => {
          console.log(result.data);
          this.position = result.data;
          console.log("This is client: request sent");
        })
        .catch(error => {
          /*eslint-disable*/
          console.error(error);
          /*eslint-enable*/
        });
    }
  }
};
</script>