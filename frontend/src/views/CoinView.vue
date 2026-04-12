<template>
  <v-container class="py-6">
    <div class="text-h5 font-weight-bold mb-6">🪙 {{ t('coin.title') }}</div>
    <v-row>
      <v-col cols="12" md="4">
        <CoinFlip />
      </v-col>
      <v-col cols="12" md="4">
        <MultiCoinFlip />
      </v-col>
      <v-col cols="12" md="4">
        <IChingConsult />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import CoinFlip from '@/components/coin/CoinFlip.vue'
import MultiCoinFlip from '@/components/coin/MultiCoinFlip.vue'
import IChingConsult from '@/components/coin/IChingConsult.vue'

const { t } = useI18n()
</script>
