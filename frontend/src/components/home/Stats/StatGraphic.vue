<template>
    <div style="padding: 0 95px">
        <v-row>
            <template v-for="n in 4">
                <v-col :key="n" style="padding: 15px">
                    <v-card class="pa-2 box">
                        <div class="justify-center">Test</div>
                    </v-card>
                </v-col>
                <v-responsive
                        v-if="n === 2"
                        :key="`width-${n}`"
                        width="100%"
                ></v-responsive>
            </template>
        </v-row>
    </div>
</template>

<script>

    export default {
        name: "StatGraphic",
        components: {

        }
    }
</script>

<style scoped>
    .box{
        padding: 15px;
        background-color: #789659 !important;
        border-radius: 12px;
        height: 95px;
    }

</style>