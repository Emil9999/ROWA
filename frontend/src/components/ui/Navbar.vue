<template>
  <v-row>
        <v-col>
           <v-btn icon>
               <v-img src="@/assets/icons/settings.svg"/>
           </v-btn>
        </v-col>
        <v-col>
            <button-home firstElement="Farm" secondElement="Home" thirdElement="Stats" @firstElement="firstFn"/>
        </v-col>
        <v-col>
         <v-btn icon>
               <v-img src="@/assets/icons/help.svg"/>
           </v-btn>
        </v-col>
      </v-row>
</template>

<script>
import ButtonHome from './ButtonHome.vue'
export default {
  components: { ButtonHome },
    methods: {
        firstFn: function(){
            console.log("hello")
        }
    }
}
</script>

<style>

</style>