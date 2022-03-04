<template>
<div>
<video alt="susteynVid" muted autoplay @ended="vidEnded" width="500" height="auto">
    <source src="../../../../frontend_old/src/assets/videos/HarvestInst.mp4" />
</video>
</div>
    
</template>

<script lang="ts">
import { defineComponent} from 'vue';

export default defineComponent({
  components: {  },
  name: 'VideoFrame',
  props:{ 
      vidType:{
          type: String,
          default: "/vid/park/"

      }
  },
  emits: ["VidEnded"],
  methods: {
    vidEnded() {
      this.$emit("VidEnded", false);
    },
  }
  
});
</script>