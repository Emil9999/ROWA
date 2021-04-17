<template>
  <div class="home">
    
    <farm-representation/>
  </div>
</template>

<script>
// @ is an alias to /src
import FarmRepresentation from '@/components/farm/FarmRepresentation.vue'

export default {
  name: 'Home',
  components: {
    
    FarmRepresentation
  }
}
</script>
