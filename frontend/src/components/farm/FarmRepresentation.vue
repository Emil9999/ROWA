<template>
  <div style="position: relative">
    <svg
      width="410"
      height="522"
      viewBox="0 0 410 522"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <rect x="190" width="30" height="494" fill="#BDBDBD" />
      <path d="M0 522L0 494H410V522H0Z" fill="#BDBDBD" />
      <rect
        width="190"
        height="7"
        transform="matrix(-1 0 0 1 190 0)"
        fill="#BDBDBD"
      />
      <rect
        width="190"
        height="7"
        transform="matrix(-1 0 0 1 410 0)"
        fill="#BDBDBD"
      />
    </svg>
    <div class="module" style="bottom: 371px; right: 0px;">
      <farm-module :id="6" :left="false"/>
    </div>
    <div class="module" style="bottom: 325px;">
      <farm-module :id="5" :left="true"/>
    </div>
    <div class="module" style="bottom: 300px; right: 0px;">
      <farm-module :id="4" :left="false"/>
    </div>
    <div class="module" style="bottom: 200px;">
      <farm-module :id="3" :left="true"/>
    </div>
    <div class="module" style="bottom: 150px; right: 0px;">
      <farm-module :id="2" :left="false"/>
    </div>
    <div class="module" style="bottom: 100px; ">
      <farm-module :id="1" :left="true"/>
    </div>
  </div>
</template>

<script>
import FarmModule from "./FarmModule.vue";

export default {
  components: { FarmModule },
  
};
</script>

<style lang="scss" scoped>
.module {
  position: absolute !important;
}
</style>
