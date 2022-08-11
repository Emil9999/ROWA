<template>
    <v-btn-toggle rounded class="shadow" borderless>
        <v-btn :style= "[farm_active==0 ? {'background-color': 'white', 'color': 'var(--v-primary-base)'}  : {'background-color': 'var(--v-primary-base)', 'color': 'white'}]"
        @click="resetFarmInfo(0)">
            Farm
        </v-btn>
        <v-btn :style= "[farm_active==1 ? {'background-color': 'white', 'color': 'var(--v-primary-base)'} : {'background-color': 'var(--v-primary-base)', 'color': 'white'}]"
               @click="resetFarmInfo(1)">
            Home
        </v-btn>
        <v-btn :style= "[farm_active==2 ? {'background-color': 'white', 'color': 'var(--v-primary-base)'} : {'background-color': 'var(--v-primary-base)', 'color': 'white'} ]"
               @click="resetFarmInfo(2)">
            Stats
        </v-btn>
    </v-btn-toggle>
</template>

<script>
    import {mapState} from "vuex"

    export default {
        name: "ButtonPill",
        computed: {
            ...mapState(["farm_active"])
        },
        methods: {

            resetFarmInfo: function(target){
                this.$store.dispatch('change_ypos_plantInfo', 260);
                this.$store.dispatch('change_ypos_statInfo', 260);
                this.$store.dispatch('change_dash_state', target)
            }
        }
    }
</script>

<style scoped>
    .shadow{box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.25);}
</style>