<template>
  <div>
    <a v-if="href" :href="href" class="button">
      <slot/>
    </a>
    <button v-else class="button">
      <slot/>
    </button>
  </div>
</template>

<script>
  export default {
    props: {
      href: {
        type: String,
        default: null
      }
    }
  }
</script>