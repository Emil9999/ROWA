<template>
    <div style="position: relative; width: 350px;height: 160px">
       
        <div style="position: absolute; bottom: 18px; width: 350px; height: 80px;">
            <div v-for="key in positions" :key="key" style="width: 15%; display: inline-flex; height: 40px; z-index: 3; align-items: center;
  justify-content: center;">
          
                <img v-if="module_plants.pos[key]" :height="calculate_img_size(module_plants.pos[key].age)" :src="getImgUrl(module_plants.type)" v-bind:style="{ 'margin-top': 40-calculate_img_size(module_plants.pos[key].age) + 'px' }">
                <div v-else style="display: inline-block; width: 34px"></div>
                
                <div v-if="key == pos" style="position: absolute; bottom: -70px;  z-index: 9;"> 
<svg width="56" height="149" viewBox="0 0 56 149" fill="none" xmlns="http://www.w3.org/2000/svg">
<path d="M28.0607 104.939C27.4749 104.354 26.5251 104.354 25.9393 104.939L16.3934 114.485C15.8076 115.071 15.8076 116.021 16.3934 116.607C16.9792 117.192 17.9289 117.192 18.5147 116.607L27 108.121L35.4853 116.607C36.0711 117.192 37.0208 117.192 37.6066 116.607C38.1924 116.021 38.1924 115.071 37.6066 114.485L28.0607 104.939ZM28.5 149L28.5 106H25.5L25.5 149H28.5Z" fill="#E3927B"/>
<line x1="1.5" y1="0.619385" x2="1.5" y2="85.8218" stroke="#E3927B" stroke-width="3"/>
<line y1="-1.5" x2="54.0781" y2="-1.5" transform="matrix(-0.999999 0.00111507 -0.000779499 -1 54.0781 85.2024)" stroke="#E3927B" stroke-width="3"/>
<line y1="-1.5" x2="54.0781" y2="-1.5" transform="matrix(-0.999999 0.00115996 -0.00074933 -1 54.0781 0)" stroke="#E3927B" stroke-width="3"/>
<line x1="54.0195" x2="54.0195" y2="85.2024" stroke="#E3927B" stroke-width="3"/>
</svg></div>
            </div>
        </div>
        <div style="position: absolute; bottom: 0">
            <svg width="350" height="48" viewBox="0 0 350 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect width="349.999" height="41" transform="matrix(-1 0 0 1 350 0)" fill="#C4C4C4"/>
                <rect width="349.999" height="7" transform="matrix(-1 0 0 1 350 41)" fill="#828282"/>
            </svg>
              
        </div>
        <div style="position: absolute; bottom: 30px; width: 340px; height: 80px;">
            <div v-for="key in positions" :key="key" style="width: 15%; display: inline-flex; height: 50px; z-index: 2; align-items: center;
  justify-content: center;">
                <span class="dot" v-if="module_plants.pos[key].harvestable" :style="[module_plants.pos[key].harvestable ? {'background-color': '#789659'} : {}]"></span>
                <span class="dot" v-if="module_plants.pos[key].harvestable == null" :style="[{'background-color': '#E3927B'}]"></span>
            </div>
        </div>
    </div>

</template>

<script>
    import {mapState, mapGetters} from "vuex"
    
    export default {
        name: "Module",
        props: ["id", "reverse", "pos"],
    
        computed: {
            ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.id]
            },
            positions: function () {
                if (this.reverse){
                    return Object.keys(this.module_plants.pos)
                }
                else {
                    return Object.keys(this.module_plants.pos).reverse()
                }
            }
        },
        methods: {
            ...mapGetters(["get_module"]),
            getImgUrl(pic) {
                return require('@/assets/cat_tree/modules/'+pic+".svg")
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
        created() {
            
        }
    }
</script>

<style scoped>
    .dot {
        bottom: 0;
        position: absolute;
        height: 10px;
        width: 10px;
        border-radius: 50%;
    }
</style>