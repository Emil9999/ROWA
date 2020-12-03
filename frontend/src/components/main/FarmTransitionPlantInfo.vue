<template>
    <div>
        <div class="moving-card" ref="interactElement"
             :style="{ transform: transformString}"
             :class="{ isAnimating:  isInteractAnimating} "
        >
            <v-row justify="center">
                <img :src="require('@/assets/main/rectangle_28.svg')" alt="line" style="padding-top: 12px">
            </v-row>
            <slot></slot>
        </div>
    </div>
</template>

<script>
    import interact from "interactjs"
    import {mapState} from "vuex"

    export default {
        name: "FarmTransitionPlantInfo",
        props: {
            yPositions: Array
        },
        data() {
            return {
                interactPosition: {
                    x: 0,
                    y: this.yPos_plantInfo
                },
                isInteractAnimating: true,
            }
        },
        methods: {
            interactSetPosition(coordinates) {
                const {x = 0, y = 0} = coordinates;
                this.interactPosition = {x, y};
            },
            stopCardAtPosition() {
                let Ylist = [...this.yPositions]
                const difference = Ylist.map(yElement => Math.abs(yElement - this.interactPosition.y))
                let i = difference.indexOf(Math.min(...difference));
                const yMovement = this.yPositions[i]
                this.interactSetPosition({x: 0, y: yMovement});
            }
        },
        computed: {
            transformString() {
                const {x, y} = this.interactPosition;
                return `translate3D(${x}px, ${y}px, 0)`;
            },
        ...mapState(["yPos_plantInfo"]),
        },
        watch: {
            
            yPos_plantInfo: function(val){
                this.interactPosition.y = val
            }
        },
        mounted() {
            const element = this.$refs.interactElement;
            interact(element).draggable({
                lockAxis: "y",
                onstart: () => {
                    this.isInteractAnimating = false;
                },
                onmove: event => {
                    const x = this.interactPosition.x + event.dx;
                    const y = this.interactPosition.y + event.dy;
                    this.interactSetPosition({x, y});
                },
                onend: () => {
                    this.isInteractAnimating = true;
                    this.stopCardAtPosition();
                }
            });
        },
        beforeDestroy() {
            interact(this.$refs.interactElement).unset();
        }
    }
</script>

<style scoped>
    .moving-card {
        top: 75%;
        position: fixed;
        background: white;
        height: 100%;
        width: 100%;
        border-radius: 60px 60px 0 0;
        box-shadow: 0px 4px 30px rgba(0, 0, 0, 0.25);

    }

    .isAnimating {
        transition: transform 0.7s;
    }
</style>