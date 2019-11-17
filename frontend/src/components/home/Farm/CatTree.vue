<template>
    <div style="position: relative; width: 410px; height: 522px" >
        <svg width="410" height="522" viewBox="0 0 410 522" fill="none" xmlns="http://www.w3.org/2000/svg">
            <rect x="190" width="30" height="494" fill="#828282"/>
            <path d="M0 522L0 494H410V522H0Z" fill="#828282"/>
            <rect width="190" height="7" transform="matrix(-1 0 0 1 190 0)" fill="#828282"/>
            <rect width="190" height="7" transform="matrix(-1 0 0 1 410 0)" fill="#828282"/>
        </svg>
        <div class="module" style="bottom: 371px">
            <Module :id="1" style=""></Module>
        </div>
        <div class="module" style="bottom: 325px; right: 0px">
            <Module :id="2" style=""></Module>
        </div>
        <div class="module" style="bottom: 247px;">
            <Module :id="3" style=""></Module>
        </div>
        <div class="module" style="bottom: 191px; right: 0">
            <Module :id="4" style=""></Module>
        </div>
        <div class="module" style="bottom: 123px;">
            <Module :id="5" style=""></Module>
        </div>
        <div class="module" style="bottom: 77px; right: 0">
            <Module :id="6" style=""></Module>
        </div>
    </div>

</template>

<script>
    import Module from "./Module";
    export default {
        name: "CatTree",
        components: {
            Module
        }
    }
</script>

<style scoped>
    .module{
        position: absolute !important;
    }
</style>