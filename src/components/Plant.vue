<template>
<div class="container">
  <div class="row">
    <h1>Plant</h1>
    <img src="../assets/logo_rowa.png">

  </div>
  
<div>
  <h2>Choose a plant type</h2>
  <!-- Get available plant types from db and generate list based on that-->
  <b-button v-on:click="postreq">Lettuce</b-button>
  <h2> Please plant in position {{position}}</h2>
</div>
</div>
</template>
<script>
import axios from 'axios';

export default {
  data () {
    return {
      position: 0
    }
  },
  methods: {
      postreq: function(){
          axios.post("http://127.0.0.1:3000/plant/greet",
          {plantType:"Lettuce"},
           "content-type: application/json").then(result => {
             console.log(result.data)
             console.log("This is client: request sent")
          }).catch(

            error => {
            /*eslint-disable*/
            console.error(error);
            /*eslint-enable*/
            })
      }
    

  }  
}
</script>