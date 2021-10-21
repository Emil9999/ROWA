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
        background-color: #F7F9F7;
        font-family: Montserrat,sans-serif;
        font-style: normal;
        user-select: none;
    }
</style>
