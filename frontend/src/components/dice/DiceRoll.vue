<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ t('dice.roll') }} d{{ sides }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.lastRoll" :anim-key="animKey">
        <div class="text-h1 mb-1">{{ diceEmoji(store.lastRoll?.value) }}</div>
        <div class="text-h4 font-weight-bold">{{ store.lastRoll?.value }}</div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-dice-6"
        @click="doRoll"
      >
        {{ t('dice.roll') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useDiceStore } from '@/stores/dice'
import ResultCard from '@/components/common/ResultCard.vue'

const props = withDefaults(defineProps<{ sides?: number }>(), { sides: 6 })

const { t } = useI18n()
const store = useDiceStore()
const animKey = ref(0)

const DICE_EMOJI: Record<number, string> = {
  1: '⚀', 2: '⚁', 3: '⚂', 4: '⚃', 5: '⚄', 6: '⚅',
}

function diceEmoji(value?: number): string {
  if (!value) return '🎲'
  return DICE_EMOJI[value] ?? '🎲'
}

async function doRoll() {
  animKey.value++
  await store.roll(props.sides)
}
</script>
