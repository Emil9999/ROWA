<template>
    <v-row>
        <v-col cols="4" style="padding-left: 30px">
            <v-btn dark fab color="white">
                <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
        </v-col>
        <v-col align-content-center cols="4">
           
        </v-col> 
        <v-col cols="4" style="float:right">
            <v-btn dark fab color="white">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
        </v-col>
    </v-row>

</template>

<script>

    export default {
        name: "FarmingTopRow",
        components:{
        
        }
    }
</script>

<style scoped>
    #logo-text{
        font-family: Montserrat, sans-serif;
        font-style: normal;
        font-weight: 600;
        font-size: 36px;
        color: var(--v-primary-base)
    }
</style>