import { reactive } from 'vue'

export const store = reactive({
  unlocked: false,
  change() {
    this.unlocked = !this.unlocked
  }
})
