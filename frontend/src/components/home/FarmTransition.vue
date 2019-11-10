<template>
    <div>
        <div class="box">
            <div class="hid-box" :class="{ up: farmInfo}" v-touch:swipe.top="farmInfo = true" v-touch:swipe.bottom="farmInfo = false">
                <h1 style="padding: 20px">CSS3 slide up</h1>
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

    .up{
        top: 20px;
    }
</style>