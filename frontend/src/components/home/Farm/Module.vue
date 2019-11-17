<template>
    <div style="position: relative; width: 190px;height: 160px">
        <div style="position: absolute; bottom: 22px; width: 190px; height: 80px;">
            <div v-for="(pos, key) in module_plants.pos" :key="key" style="max-width: 20%; min-width:20px; display: inline-block">
                <img v-if="pos" :height="calculate_img_size(pos.age)" :src="getImgUrl(module_plants.type)">
                <div v-else style="display: inline-block; width: 26px"></div>
            </div>
        </div>
        <div style="position: absolute; bottom: 0">
            <svg width="190" height="48" viewBox="0 0 190 48" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect width="189.999" height="41" transform="matrix(-1 0 0 1 190 0)" fill="#C4C4C4"/>
                <rect width="189.999" height="7" transform="matrix(-1 0 0 1 190 41)" fill="#828282"/>
            </svg>
        </div>
    </div>

</template>

<script>
    import {mapState, mapGetters} from "vuex"
    
    export default {
        name: "Module",
        props: ["id"],
        computed: {
            ...mapState(["farm_info"]),
            module_plants: function () {
                return this.farm_info[this.$props.id]
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
                    size = 10
                }
                else if (days > 36){
                    size = 50
                }
                else {
                    size = days/42 * 50
                }
                return size
            },
        },
        created() {
            console.log(this.module_plants)
        }
    }
</script>

<style scoped>

</style>