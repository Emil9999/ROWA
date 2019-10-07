<template>
<div>
<h1>I HATE VUE</h1>
<b-button v-on:click="postreq">Hello</b-button>
</div>
</template>
<script>
import axios from 'axios';

export default {
  methods: {
      postreq: function(){
          axios.post("http://127.0.0.1:3000/plant/greet",
          {name:"Joe",email:"joe@labstack"},
           "content-type: application/json").then(result => {
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