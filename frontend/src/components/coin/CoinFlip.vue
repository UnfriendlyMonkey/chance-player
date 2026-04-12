<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ t('coin.flip') }}</v-card-title>
    <v-card-text>
      <ResultCard :show="!!store.lastFlip" :anim-key="animKey">
        <div class="text-h2 mb-2">{{ store.lastFlip?.side === 'heads' ? '🪙' : '🔘' }}</div>
        <div class="text-h5 font-weight-bold">
          {{ store.lastFlip ? t(`coin.result.${store.lastFlip.side}`) : '' }}
        </div>
      </ResultCard>
      <div v-if="store.error" class="text-error mt-2">{{ t('common.error') }}</div>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-circle-half-full"
        @click="doFlip"
      >
        {{ t('coin.flip') }}
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
  await store.flip()
}
</script>
