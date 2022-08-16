/<template>
  <div
    :class="{
      'opacity-50':
        harvestableSpots.length + plantableSpots.length == 0 && !showUnavailable,
    }"
    class="w-72 h-32 flex-col flex"
  >
    <div
      :class="[
        reverseModule,
        '-mx-' + (7 - count) * 2 +
          ' basis-1/2 inline-flex justify-between items-end',
      ]"
    >
      <div
        v-for="position in count"
        :key="position"
        :style="'width:' + 72 / count / 4 + 'rem;'"
        class="h-12 text-center mx-auto flex items-end justify-cen5ter"
      >
        <img
          v-if="findPlant(position)"
          :class="[' mx-auto']"
          :width="
            plantWidth(
              findPlant(position)?.age,
              findPlant(position)?.growth_time
            )
          "
          :src="
            require('../../../assets/img/plant_svg/' +
              cImage(findPlant(position)?.plant_type))
          "
        />
        <img
          v-if="!findPlant(position)"
          :class="[emptySpaceClass, ' mx-auto']"
          :width="plantWidth(0, 70)"
          :src="require('../../../assets/img/plant_svg/default.svg')"
        />
      </div>
    </div>

    <div class="border-t-4 border-grey">
      <div :class="reverseModule" class="h-full basis-1/4 flex justify-between">
        <div
          :class="{
            invisible: !(
              harvestableSpots.find(e => e.position == i) || plantableSpots.find(e => e.position == i)
            ),
          }"
          v-for="i in count"
          :key="i"
          class="
            rounded-full
            flex flex-none
            justify-center
            content-start
            -mb-7
            bg-almostwhite
            w-8
          "
        >
          <div
            v-if="harvestableSpots.find(e => e.position == i)"
            class="bg-green mt-0.5 rounded-full h-7 w-7"
          >
            <CheckIcon class="text-white" />
          </div>
          <div
            v-if="plantableSpots.find(e => e.position == i)"
            class="bg-brownred mt-0.5 rounded-full h-7 w-7"
          >
            <ArrowSmDownIcon class="text-white" />
          </div>
        </div>
      </div>

      <div
        :class="reverseModule"
        class="
          flex-none
          rounded-full
          basis-1/4
          flex
          items-baseline
          h-8
          justify-between
          bg-almostwhite
        "
      >
        <div class="h-1/12 w-1/12"></div>
        <div class="p-small-grey">{{ moduleText }}</div>
        <div
          class="
            text-center
            self-center
            bg-accentwhite
            ml-0.5
            rounded-full
            h-7
            w-7
            p-green-small
          "
        >
          {{ moduleNumber }}
        </div>
      </div>
    </div>
  </div>
</template>
/<script lang="ts">
import { computed, defineComponent, ref, inject, PropType } from "vue";
import { CheckIcon, ArrowSmDownIcon } from "@heroicons/vue/solid";
import getPlantsInModule from "../../../composables/use_getPlantInModule";
import { checkImage } from "../../../composables/use_imgChecker";
import FarmablePlant from '../../../types/FarmablePlant'
import findModuleGroup from "../../../composables/use_findModuleGroup";

export default defineComponent({
  components: { CheckIcon, ArrowSmDownIcon },
  setup(props) {
    // Array plantable harvestable pos

    const clamp = (num: number, min: number, max: number) =>
      Math.min(Math.max(num, min), max);

    const plantWidth = (age: number, gTime: number) => {
      gTime = gTime == 0 ? 1 : gTime;
      let wnumber = 10 + clamp(60 * (age / gTime), 0, 60);
      return wnumber;
    };

    

    const { group, findGroup } = findModuleGroup();
    findGroup(props.moduleNumber);
    const showUnavailable = inject("showunavailable", false);
    const {
      modulePlants: plantsInModule,
      loadModulePlants,
      plantcountInModule: count,
    } = getPlantsInModule(props.moduleNumber);
    const emptySpaceClass = ref("filter grayscale opacity-75");
    const { cImage } = checkImage("svg");
    const harvestable = ref(props.harvestableSpots);
    const plantable = ref(props.plantableSpots);
    const reverseModule = computed(() => ({
      "flex-row-reverse": props.reverse,
    }));
    loadModulePlants();

    const findPlant = (position: number) => {
      return plantsInModule.value.find((e) => e.position == position);
    };

   

    const moduleText = computed(() => {
      if (plantsInModule.value.length == 0) {
        return "Empty";
      }
      if (props.harvestableSpots.length + props.harvestableSpots.length == 0 && !showUnavailable) 
      {
        return "Unavailable";
      } else 
      {
        if (group.value == "herb") {
            return "Herb";
        } 
        else{
            return plantsInModule.value.find((e) => e.plant_type !== "")
                ?.plant_type;
            }
      }
    });

    return {
      count,
      plantWidth,
      cImage,
      moduleText,
      emptySpaceClass,
      plantsInModule,
      reverseModule,
      harvestable,
      plantable,
      showUnavailable,
      findPlant,
    };
  },
  props: {
    moduleNumber: {
      type: Number,
      default: 1,
    },
    reverse: {
      type: Boolean,
      default: false,
      required: false,
    },
    plantableSpots: {
      type: Array as PropType<Array<FarmablePlant>>,
      required: true,
      default: () => [],
    },
    harvestableSpots: {
      type: Array as PropType<Array<FarmablePlant>>,
      required: true,
      default: () => [],
      
    },
  },
  methods: {},
});
</script>
<style>
text {
  font-family: Mulish;
  font-style: normal;
  font-weight: 600;
  font-size: 36px;

  /* UI COLORS/Grey Dark */

  fill: #828282;
}
</style>