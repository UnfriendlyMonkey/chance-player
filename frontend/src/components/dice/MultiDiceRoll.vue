<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ label }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.lastMulti" :anim-key="animKey">
        <div class="d-flex justify-center gap-3 text-h2 mb-2">
          <span v-for="(roll, i) in store.lastMulti?.rolls" :key="i">
            {{ DICE_EMOJI[roll.value] ?? '🎲' }}
          </span>
        </div>
        <div class="text-h6 text-medium-emphasis">
          {{ t('dice.sum') }}: <strong>{{ store.lastMulti?.sum }}</strong>
        </div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-dice-multiple"
        @click="doRoll"
      >
        {{ label }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useDiceStore } from '@/stores/dice'
import ResultCard from '@/components/common/ResultCard.vue'

const props = withDefaults(
  defineProps<{ count?: number; sides?: number }>(),
  { count: 2, sides: 6 },
)

const { t } = useI18n()
const store = useDiceStore()
const animKey = ref(0)

const DICE_EMOJI: Record<number, string> = {
  1: '⚀', 2: '⚁', 3: '⚂', 4: '⚃', 5: '⚄', 6: '⚅',
}

const label = computed(() => t(`dice.roll${props.count}`) || `Roll ${props.count}d${props.sides}`)

async function doRoll() {
  animKey.value++
  await store.rollMultiple(props.count, props.sides)
}
</script>
