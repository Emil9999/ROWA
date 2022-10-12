/<template>
  <div v-if="!detailEnter" class="centered-div">
        <div class="h-green-big my-8">New Plant Type</div>
        <div class="grid grid-cols-2 gap-10">
            <div class="grid grid-rows-3 gap-2 h-green-small">
                <div>Name:</div>
                <div>Grow Time:</div>
                <div>Type:</div>
            </div>
            <div  class="grid grid-rows-3 gap-2">
                <div  v-for="(item, key) in newType" :key="item">
                    <div class="flex" v-if="key == 'group'"> 
                        <div @click="newType.group = 'herb'" :class="[item == 'herb' ?'btn-admin-green':'btn-admin-white']">Herb</div>
                        <div @click="newType.group = 'lettuce'" :class="[item == 'lettuce' ?'btn-admin-green':'btn-admin-white']">Lettuce</div>
                    </div>
                    <div class="btn-admin-white" v-else @click="detailEnter = true, selectedKey = key, tempNumber = newType.growth_time">{{item == '' ? '-' : item}}</div>
                </div>
            </div>
        </div>
        <div v-if="!allFilled"  @click="$emit('saveNewType', newType.plant_type, newType.group)" class="btn-big-green">Save</div>
  </div>
  <div v-else class="centered-div">
        <div class="rounded-full p-3 left-0 bg-green" @click="detailEnter = false">Back</div>
        <div v-if="selectedKey == 'growth_time'">
             <div class="h-green-small">{{tempNumber}}</div>
             <numberPad :altText="'Save'" @clickedAlt="newType.growth_time = tempNumber, detailEnter = false" @backClicked="tempNumber = tempNumber.slice(0, -1)" @numberClicked="n => tempNumber += n"></numberPad>
        </div>

        <div v-else> 
            <keyboard @saveName="addName"/> 
        </div>
  </div>
</template>

<script lang="ts">
import { computed, defineComponent, ref } from 'vue'
import keyboard from '../farming/atoms/keyboard.vue'
import numberPad from './atoms/numberPad.vue'



export default defineComponent({
    components:{keyboard, numberPad},
    setup() {
        
        const allFilled = computed(() => {
            
            return Object.values(newType.value).includes(0 || '')
            
        })

        const tempNumber = ref('')
        const selectedKey = ref<string | null>()
        const detailEnter = ref(false)
        function addName(name:string){
                newType.value.plant_type = name
                detailEnter.value = false
        }
        const newType = ref({plant_type: '', growth_time: '', group: ''})
        return{detailEnter, newType, selectedKey, addName, allFilled, tempNumber}
    },
})
</script>
