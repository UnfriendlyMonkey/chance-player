<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">📖 {{ t('coin.iching') }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.iching" :anim-key="animKey">
        <div
          v-for="(line, i) in store.iching?.lines"
          :key="i"
          class="d-flex align-center gap-2 mb-2"
        >
          <span class="text-caption text-medium-emphasis" style="width: 3rem">
            {{ t('coin.line') }} {{ i + 1 }}
          </span>
          <span class="iching-line font-weight-bold" :class="lineClass(line.line_type)">
            {{ lineSymbol(line.line_type) }}
          </span>
          <div class="d-flex gap-1">
            <v-chip
              v-for="(side, j) in line.coins"
              :key="j"
              size="x-small"
              :color="side === 'heads' ? 'primary' : 'secondary'"
              variant="tonal"
            >
              {{ side === 'heads' ? 'H' : 'T' }}
            </v-chip>
          </div>
          <span class="text-caption text-medium-emphasis">({{ line.sum }})</span>
          <v-chip size="x-small" :color="isChanging(line.line_type) ? 'warning' : 'default'" variant="tonal">
            {{ t(`coin.lineTypes.${line.line_type}`) }}
          </v-chip>
        </div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-book-open"
        @click="doConsult"
      >
        {{ t('coin.consult') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCoinStore } from '@/stores/coin'
import ResultCard from '@/components/common/ResultCard.vue'
import type { IChingLineType } from '@/api/types'

const { t } = useI18n()
const store = useCoinStore()
const animKey = ref(0)

async function doConsult() {
  animKey.value++
  await store.iChing()
}

function lineSymbol(type: IChingLineType): string {
  return type === 'old_yin' || type === 'young_yin' ? '— —' : '———'
}

function lineClass(type: IChingLineType): string {
  return isChanging(type) ? 'text-warning' : ''
}

function isChanging(type: IChingLineType): boolean {
  return type === 'old_yin' || type === 'old_yang'
}
</script>

<style scoped>
.iching-line {
  font-family: monospace;
  font-size: 1.1rem;
  min-width: 3rem;
  letter-spacing: 0.1em;
}
</style>
