<template>
    <v-row>
        <v-col cols="4" style="padding-left: 20px">
            <v-btn icon x-large :to="{name:'AdminMenu'}">
                <v-icon>$vuetify.icons.adminSettings</v-icon>
            </v-btn>
        </v-col>
        <v-col cols="5" justify-self="center">
            <ButtonPill></ButtonPill>
        </v-col>
        <v-col cols="3" id="logo-text" style="text-align: right">
            SUSTEYN
        </v-col>
    </v-row>

</template>

<script>
    import ButtonPill from "./ButtonPill";

    export default {
        name: "HomeTopRow",
        components:{
            ButtonPill
        }
    }
</script>

<style scoped>
    #logo-text{
        font-family: Montserrat, sans-serif;
        font-style: normal;
        font-weight: 600;
        font-size: 32px;
        color: var(--v-primary-base)
    }
</style>