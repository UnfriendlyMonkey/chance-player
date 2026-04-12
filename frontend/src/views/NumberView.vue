<template>
  <v-container class="py-6">
    <div class="text-h5 font-weight-bold mb-6">🎯 {{ t('number.title') }}</div>
    <v-row>
      <v-col cols="12" sm="6" md="4">
        <RandomNumber />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import RandomNumber from '@/components/number/RandomNumber.vue'

const { t } = useI18n()
</script>
