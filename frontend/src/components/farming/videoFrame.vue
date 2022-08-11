<template>
<div>
  <div v-if="vidURL">
  <video alt="susteynVid" muted autoplay @ended="this.$emit('VidEnded', false)" width="500" height="auto">
      <source src="../../../../frontend_old/src/assets/videos/HarvestInst.mp4" />
  </video>
  </div>
  <div v-if="!vidURL">
    <video alt="susteynVid" muted autoplay @ended="this.$emit('VidEnded', false)" width="500" height="auto">
      <source src="../../../../frontend_old/src/assets/videos/HarvestInst.mp4" />
  </video>
  </div>
</div>
    
</template>

<script lang="ts">
import { computed, defineComponent, ref} from 'vue';

export default defineComponent({
  components: {  },
  name: 'VideoFrame',
  props:{ 
      vidType:{
          type: String,
          default: "hsalad"

      }
  },
  setup(props){
    const vidURL = () => {
      if(props.vidType == 'hsalad'){
          return true} else{return false}
    }


    return {vidURL}
  }
  
});
</script>