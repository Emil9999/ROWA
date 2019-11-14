<template>
    <div>
        <div class="box">
            <div class="moving-card" :class="{ up: farmInfo}" @click="farmInfo = !farmInfo">
                <h1>CSS3 slide up</h1>
                <p>This is a quick demo of slide-up effect using CSS animation when hover the box. No JS required!</p>
            </div>
        </div>
    </div>

</template>

<script>
    export default {
        name: "demo",
        data() {
            return {
                farmInfo: false
            }
        }
    }
</script>

<style scoped>
    h1 {
        margin: 0
    }
    .box {
        height: 200px;
        width: 100%;
        bottom: 0;
        overflow: hidden;
        border: 1px solid red;
        background: #ff0;
        position: absolute;
    }
    .moving-card {
        top: 100px;
        position: relative;
        transition: all .3s ease-out;
        background: white;
        height: 100%;
    }

    .up{
        top: 0;
    }
</style>