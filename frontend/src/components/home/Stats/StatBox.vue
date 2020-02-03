<template>
    <div class="info-box">
        <v-row justify="center">
            <h2>{{heading}}</h2>
        </v-row>
        <v-row justify="center" style="padding-top: 10px">
            <VueSpeedometer v-if="type === 'temperature'" :value="value" :height="130" :width="210" :segments="5"
                            :segmentColors='["tomato", "gold", "limegreen", "gold", "tomato"]'
                            :minValue="0" :maxValue="40" textColor="white"
                            currentValueText="${value} °C"/>
            <VueSpeedometer v-if="type ==='light'" :value="value" :height="130" :width="210" :segments="1000"
                            :minValue="0" :maxValue="1000" :maxSegmentLabels="5"
                            currentValueText="${value} Lux" textColor="white"/>
        </v-row>
    </div>
</template>

<script>
    import VueSpeedometer from "vue-speedometer"

    export default {
        name: "StatBox",
        components: {
            VueSpeedometer
        },
        props: {
            heading: String,
            value: Number,
            type: String
        },
    }
</script>

<style scoped>
    .info-box {
        background: #789659;
        border-radius: 10px;
        box-shadow: 0px 4px 4px rgba(0, 0, 0, 0.25);
        padding-bottom: 10px;
        color: white;
    }

    h2 {
        padding-top: 10px;
        font-size: 16px;
    }
</style>