<template>
    <div style="position: relative; width: 190px;height: 160px">
        <div style="position: absolute; bottom: 18px; width: 210px; height: 80px;">
            <div v-for="key in positions" :key="key" style="width: 15%; display: inline-flex; height: 40px; z-index: 3; align-items: center;
  justify-content: center;">
                <img v-if="module_plants.pos[key]" :height="calculate_img_size(module_plants.pos[key].age)" :src="getImgUrl(module_plants.type)" v-bind:style="{ 'margin-top': 40-calculate_img_size(module_plants.pos[key].age) + 'px' }">
                <div v-else style="display: inline-block; width: 34px"></div>
            </div>
        </div>
        <div style="position: absolute; bottom: 0">
            <svg width="190" height="48" viewBox="0 0 190 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect width="189.999" height="41" transform="matrix(-1 0 0 1 190 0)" fill="#C4C4C4"/>
                <rect width="189.999" height="7" transform="matrix(-1 0 0 1 190 41)" fill="#828282"/>
            </svg>
        </div>
        <div style="position: absolute; bottom: 30px; width: 210px; height: 80px;">
            <div v-for="key in positions" :key="key" style="width: 15%; display: inline-flex; height: 50px; z-index: 2; align-items: center;
  justify-content: center;">
                <span class="dot" v-if="module_plants.pos[key].harvestable" :style="[module_plants.pos[key].harvestable ? {'background-color': '#789659'} : {}]"></span>
                <span class="dot" v-if="module_plants.pos[key].harvestable == null && bplantable(key)" :style="[{'background-color': '#E3927B'}]"></span>
            </div>
        </div>
    </div>

</template>

<script>
    import {mapState, mapGetters} from "vuex"
    import axios from "axios"
    
    export default {
        name: "Module",
        data()  {
            return{
                plantamodule: null,
              bplant: true,
              reddots: [2,2,2,2,2,2],
            }
        },

        props: ["id", "reverse"],
        
    
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
            },
            upperCasePlant: function () {
                return this.module_plants.type.replace(/\b\w/g, l => l.toUpperCase())
            },
    
        },
        methods: {
            ...mapGetters(["get_module"]),
            getImgUrl(pic) {
                return require('@/assets/cat_tree/modules/'+pic+".svg")
            },
            makeRedDots: function(){
                for(var key in this.positions){
                                if(this.positions[key].harvestable == null && this.positions.find(o => o.age === 0) && this.positions.find(o => o.harvestable === null) && this.plantamodule.find(o => o.available_plants === this.id)){
                            this.reddots[key] = 1
                                } else {
                                    this.reddots[key] = 0
                                }
                            }
            },
            bplantable: function (run) {
                
               
                var rev = this.reddots.reverse();
              
                var firstapp = 5-(rev.findIndex(o => o === 1))
                
               if(firstapp == run){
                   console.log(firstapp)
                   return true
               } else {
                   return false
               }
                },
            getPlantableModules: function () {
                axios.get("http://127.0.0.1:3000/dashboard/plantable-modules")
                    .then(result => {
                        this.plantamodule = result.data
                    })
                    .catch(error => {
                        console.log(error)
                    })
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
            this.getPlantableModules()
           
        },
        mounted(){
            this.makeRedDots()
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