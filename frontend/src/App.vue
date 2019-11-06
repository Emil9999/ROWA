<template>
    <v-app id="app">
        <router-view/>
    </v-app>
</template>

<script>
    export default {
        name: 'App',
        components: {
        },
        data: () => ({
            //
        }),
    };
</script>

<style>
    @import url('https://fonts.googleapis.com/css?family=Montserrat&display=swap');

    #app {
        background-color: #F4F0E6;
    }
</style>
