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
            <VueSpeedometer v-if="type ==='e-temperature'" :value="value" :height="130" :width="210" :segments="500"
                            startColor="limegreen" endColor="tomato"
                            :minValue="10" :maxValue="70" :maxSegmentLabels="5"
                            currentValueText="${value} °C" textColor="white"/>
            <VueSpeedometer v-if="type ==='ph'" :value="value" :height="130" :width="210" :segments="14"
                            :minValue="0" :maxValue="14" :maxSegmentLabels="5"
                            currentValueText="${value} pH" textColor="white"/>
            <VueSpeedometer v-if="type ==='humidity'" :value="value" :height="130" :width="210" :segments="30"
                            :minValue="20" :maxValue="60" :maxSegmentLabels="5"
                            currentValueText="${value} %" textColor="white"/>
            <VueSpeedometer v-if="type ==='waterLevel'" :value="value" :height="130" :width="210" :segments="500"
                            startColor="white" endColor="blue"
                            :minValue="0" :maxValue="100" :maxSegmentLabels="5"
                            currentValueText="${value} %" textColor="white"/>
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