<template>
    <v-btn-toggle rounded class="shadow" borderless>
        <v-btn :style= "[farm_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]"
        @click="change_dash_state">
            Farm
        </v-btn>
        <v-btn :style= "[!farm_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]"
               @click="change_dash_state">
            Stats
        </v-btn>
    </v-btn-toggle>
</template>

<script>
    import {mapState, mapActions} from "vuex"

    export default {
        name: "ButtonPill",
        computed: {
            ...mapState(["farm_active"])
        },
        methods: {
            ...mapActions(["change_dash_state"])
        }
    }
</script>

<style scoped>
    .shadow{box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.25);}
</style>