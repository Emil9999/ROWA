<template>
    
        <div style="position: absolute; bottom: 18px; width: 210px; height: 80px;">
            <div v-for="key in positions" :key="key" style="width: 15%; display: inline-flex; height: 40px; z-index: 3; align-items: center;
  justify-content: center;">
                <img v-if="module_plants.pos[key]" :height="calculate_img_size(module_plants.pos[key].age)" :src="getImgUrl(module_plants.type)" v-bind:style="{ 'margin-top': 40-calculate_img_size(module_plants.pos[key].age) + 'px' }">
                <div v-else style="display: inline-block; width: 34px"></div>
            </div>
            <svg width="190" height="48" viewBox="0 0 190 48" fill="none" xmlns="http://www.w3.org/2000/svg">
            <rect width="189.999" height="41" transform="matrix(-1 0 0 1 190 0)" fill="#C4C4C4"/>
                <rect width="189.999" height="7" transform="matrix(-1 0 0 1 190 41)" fill="#BDBDBD"/>
            </svg>

    </div>
</template>

<script>
import {mapState, mapGetters} from "vuex"
    export default {
         props: ["id", "right"],
         methods: {
             ...mapGetters(["get_module"]),
            
            getImgUrl(pic) {
                return require('@/assets/farm/plants/'+pic+".svg")
            },
            
            calculate_img_size(days){
                let size
                if (days < 7) {
                    size = 8
                }
                else if (days > 36){
                    size = 40
                }
                else {
                    size = days/42 * 40
                }
                return size
            },
        },

        computed: {
    ...mapState(["farmInfo", "modulePlantSpots"]),
    module_plants: function() {
      return this.farmInfo[this.$props.id];
    },
    
    positions: function() {
      if (this.right) {
        return Object.keys(this.module_plants.pos);
      } else {
        return Object.keys(this.module_plants.pos).reverse();
      }
    },
    upperCasePlant: function() {
      return this.module_plants.type.replace(/\b\w/g, (l) => l.toUpperCase());
    },
  },
    }
</script>

<style lang="scss" scoped>

</style>