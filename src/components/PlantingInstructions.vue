<template>
    <div class="container">
        <div class="row">
            <h1>How to plant</h1>
        </div>
        <div class="row">
            <p>Step one etc</p>
        </div>
        <div class="row">
           <router-link to="/">
        <button type="button" class="btn btn-primary" v-on:click="addPlant">Done</button>
      </router-link>
        </div>
    </div>
</template>


<script>
import { eventBus } from '../main';
export default {
    methods: {
        addPlant: function(){
        }
    },
    created(){
        eventBus.$on('sendPosition', (position) => {
            console.log("Hello")
            console.log(position)
        })
    }
}
</script>

