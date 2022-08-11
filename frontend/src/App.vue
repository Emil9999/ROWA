<template>
<div class="bg-accentwhite-light overflow-hidden h-screen w-screen">
  <router-view :key="$route.fullPath"/>
</div>

</template>

<script lang="ts">
import { defineComponent} from 'vue';

export default defineComponent({
  components: {  },
  name: 'App',
  
});
</script>
