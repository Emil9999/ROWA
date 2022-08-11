<template>
<div>
<video alt="susteynVid" muted autoplay @ended="this.$emit('VidEnded', false)" width="500" height="auto">
    <source :src="require('../../../../frontend_old/src/assets/videos/'+vidURL()+'.mp4')" />
</video>
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
    const deineMuddastinkt = 'HarvestInst'
    const vidURL = () => {
      if(props.vidType == 'hsalad'){
          return 'HarvestInst'
      } else {
         return 'PlantInst'}
    } 


    return {vidURL, deineMuddastinkt}
  }
  
});
</script>