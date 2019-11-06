<template>
    <v-btn-toggle rounded class="shadow" borderless>
        <v-btn color="primary" class="white--text">
            Farm
        </v-btn>
        <v-btn color="white" class="primary--text">
            Stats
        </v-btn>
    </v-btn-toggle>
</template>

<script>
    export default {
        name: "ButtonPill"
    }
</script>

<style scoped>
    .shadow{box-shadow: 0px 4px 20px rgba(0, 0, 0, 0.25);}
</style>