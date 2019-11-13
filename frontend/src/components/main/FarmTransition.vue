<template>
    <div>
        <div class="box">
            <div class="hid-box" :class="{ up: farmInfo}"
                 v-touch:tap="change_info">
                <v-row justify="center">
                    <img v-if="!farmInfo" :src="require('@/assets/main/rectangle_28.svg')" alt="line" style="padding-top: 12px">
                    <img v-else :src="require('@/assets/main/down.svg')" alt="line" style="padding-top: 12px">
                </v-row>
                <slot></slot>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: "FarmTransition",
        data() {
            return {
                farmInfo: false
            }
        },
        methods: {
            change_info: function () {
                this.farmInfo = !this.farmInfo
            }
        }
    }
</script>

<style scoped>
    .box {
        height: 430px;
        width: 100%;
        bottom: 0;
        overflow: hidden;
        position: absolute;
    }

    .hid-box {
        top: 150px;
        position: relative;
        transition: all .3s ease-out;
        background: white;
        height: 100%;
        border-radius: 60px 60px 0 0;
        box-shadow: 0px 4px 30px rgba(0, 0, 0, 0.25);
    }

    .up {
        top: 20px;
    }
</style>