<template>
    <FarmTransition>
        <div>
            <v-row justify="center">
                <h2>Farming</h2>
            </v-row>
            <v-row justify="center">
                <p>Plant and Harvest your very own plant from this farm for your lunch, or to take home.</p>
            </v-row>
            <v-row justify="center" style="margin-top: 40px">
                <v-btn id="button" class="text-capitalize" rounded color="accent" height="75" width="360">
                    Start Farming Now
                    <v-icon right dark>mdi-arrow-right</v-icon>
                </v-btn>
            </v-row>
        </div>
    </FarmTransition>
</template>

<script>
    import FarmTransition from "../main/FarmTransition";
    export default {
        name: "FarmInfo",
        components: {
            FarmTransition
        }
    }
</script>

<style scoped>
    h2{
        font-weight: 600;
        font-size: 36px;
        line-height: 44px;
        padding-top: 10px;
        color: var(--v-primary-base);
    }

    p{
        text-align: center;
        width: 410px;
        color: var(--v-primary-base);
        padding-top: 10px;
        margin: 0;
    }

    #button{
        font-style: normal;
        font-weight: bold;
        font-size: 24px;
        line-height: 29px;
        border-radius: 47px;
    }
</style>