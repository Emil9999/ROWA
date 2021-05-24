<template>
  <v-row>
        <v-col>
           <v-btn icon>
               <v-img max-height="50" max-width="50"  src="@/assets/icons/settings.svg"/>
           </v-btn>
        </v-col>
        <v-col>
            <button-home firstElement="Farm" secondElement="Home" thirdElement="Stats" 
            @first="$router.push({name: 'Farm'})"
            @second="$router.push({name: 'Home'})"
            @third="$router.push({name: 'Stats'})"
            
            />
        </v-col>
        <v-col>
         <v-btn icon>
               <v-img  max-height="90" max-width="90" src="@/assets/icons/help.svg"/>
           </v-btn>
        </v-col>
      </v-row>
</template>

<script>
import ButtonHome from './ButtonHome.vue'
export default {
  components: { ButtonHome },
    methods: {
    }
}
</script>

<style>

</style>