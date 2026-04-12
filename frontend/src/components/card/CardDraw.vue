<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ t('card.draw') }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.lastDraw" :anim-key="animKey">
        <div class="playing-card mx-auto" :class="cardColor">
          <div class="card-corner top-left">
            <div class="card-value">{{ store.lastDraw?.value }}</div>
            <div class="card-suit">{{ suitEmoji }}</div>
          </div>
          <div class="card-center text-h2">{{ suitEmoji }}</div>
          <div class="card-corner bottom-right">
            <div class="card-value">{{ store.lastDraw?.value }}</div>
            <div class="card-suit">{{ suitEmoji }}</div>
          </div>
        </div>
        <div class="text-body-1 mt-2">
          {{ store.lastDraw ? `${t(`card.suits.${store.lastDraw.suit}`)} ${store.lastDraw.value}` : '' }}
        </div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-cards"
        @click="doDraw"
      >
        {{ t('card.draw') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCardStore } from '@/stores/card'
import ResultCard from '@/components/common/ResultCard.vue'
import type { Suit } from '@/api/types'

const { t } = useI18n()
const store = useCardStore()
const animKey = ref(0)

const SUIT_EMOJI: Record<Suit, string> = {
  hearts: '♥',
  diamonds: '♦',
  spades: '♠',
  clubs: '♣',
}

const suitEmoji = computed(() =>
  store.lastDraw ? SUIT_EMOJI[store.lastDraw.suit] : '',
)

const cardColor = computed(() => {
  if (!store.lastDraw) return ''
  return store.lastDraw.suit === 'hearts' || store.lastDraw.suit === 'diamonds'
    ? 'card-red'
    : 'card-black'
})

async function doDraw() {
  animKey.value++
  await store.draw()
}
</script>

<style scoped>
.playing-card {
  width: 100px;
  height: 140px;
  border: 2px solid currentColor;
  border-radius: 8px;
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
  background: white;
}
.card-corner {
  position: absolute;
  display: flex;
  flex-direction: column;
  align-items: center;
  font-size: 0.75rem;
  line-height: 1;
}
.top-left { top: 4px; left: 6px; }
.bottom-right { bottom: 4px; right: 6px; transform: rotate(180deg); }
.card-center { font-size: 2rem; }
.card-red { color: #c62828; }
.card-black { color: #212121; }
</style>
