<template>
  <v-container class="py-6">
    <div class="text-h5 font-weight-bold mb-6">🎲 {{ t('dice.title') }}</div>
    <v-row>
      <v-col cols="12" sm="6" md="4">
        <DiceRoll :sides="6" />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <MultiDiceRoll :count="2" :sides="6" />
      </v-col>
      <v-col cols="12" sm="6" md="4">
        <MultiDiceRoll :count="3" :sides="6" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'
import DiceRoll from '@/components/dice/DiceRoll.vue'
import MultiDiceRoll from '@/components/dice/MultiDiceRoll.vue'

const { t } = useI18n()
</script>
