<template>
  <Transition name="result-pop" mode="out-in">
    <v-card
      v-if="show"
      :key="animKey"
      class="result-card text-center pa-6"
      elevation="2"
      rounded="xl"
    >
      <slot />
    </v-card>
  </Transition>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue'

const props = defineProps<{
  show: boolean
  // pass a value that changes on each new result to re-trigger animation
  animKey?: string | number
}>()

// force re-mount of card on animKey change for the transition to fire
const animKey = ref(props.animKey)
watch(() => props.animKey, (v) => { animKey.value = v })
</script>

<style scoped>
.result-pop-enter-active {
  animation: popIn 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
}
.result-pop-leave-active {
  animation: popOut 0.15s ease-in;
}
@keyframes popIn {
  from { opacity: 0; transform: scale(0.8); }
  to   { opacity: 1; transform: scale(1); }
}
@keyframes popOut {
  from { opacity: 1; transform: scale(1); }
  to   { opacity: 0; transform: scale(0.9); }
}
</style>
