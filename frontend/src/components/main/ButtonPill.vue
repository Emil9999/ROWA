<template>
    <v-btn-toggle rounded class="shadow" borderless>
        <v-btn :style= "[farm_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]">
            Farm
        </v-btn>
        <v-btn :style= "[!farm_active ? {'background-color': 'var(--v-primary-base)', 'color': 'white'} : {'background-color': 'white', 'color': 'var(--v-primary-base)'}]">
            Stats
        </v-btn>
    </v-btn-toggle>
</template>

<script>
    export default {
        name: "ButtonPill",
        props: {
            farm_active: {
                type: Boolean,
                default: true
            }
        }
    }
</script>

<style scoped>
    .shadow{box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.25);}
</style>