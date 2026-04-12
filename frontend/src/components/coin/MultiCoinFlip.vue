<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ t('coin.flip3') }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.lastMulti" :anim-key="animKey">
        <div class="d-flex justify-center gap-4 text-h3 mb-2">
          <span v-for="(side, i) in store.lastMulti?.sides" :key="i">
            {{ side === 'heads' ? '🪙' : '🔘' }}
          </span>
        </div>
        <div class="d-flex justify-center gap-4">
          <v-chip
            v-for="(side, i) in store.lastMulti?.sides"
            :key="i"
            :color="side === 'heads' ? 'primary' : 'secondary'"
            variant="tonal"
            size="small"
          >
            {{ t(`coin.result.${side}`) }}
          </v-chip>
        </div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-circle-multiple"
        @click="doFlip"
      >
        {{ t('coin.flip3') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useCoinStore } from '@/stores/coin'
import ResultCard from '@/components/common/ResultCard.vue'

const { t } = useI18n()
const store = useCoinStore()
const animKey = ref(0)

async function doFlip() {
  animKey.value++
  await store.flipMultiple(3)
}
</script>
