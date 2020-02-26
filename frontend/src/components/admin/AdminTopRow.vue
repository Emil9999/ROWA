<template>
    <v-row align-center>
        <v-col cols="2" style="padding-left: 30px">
             <router-link :to="{}">
            <v-btn dark fab color="white">
                <v-icon color="primary">mdi-arrow-left</v-icon>
            </v-btn>
             </router-link>
        </v-col>
        <v-col align="center"  align-self="center" cols="8">
             <h3 style="color:#828282" v-text="headtext"></h3>
        </v-col>
        <v-col cols="2" id="logo-text" style="text-align: right; padding-right: 30px">
             <router-link :to="{name:'Home'}">
            <v-btn dark fab color="white">
                <v-icon color="primary">mdi-close</v-icon>
            </v-btn>
            </router-link>
        </v-col>
    </v-row>

</template>

<script>
    
    export default {
        name: "AdminTopRow",
        components:{
            
        },
        props:{
                headtext: String
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