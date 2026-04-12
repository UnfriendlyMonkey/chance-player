<template>
  <v-container class="py-8">
    <div class="text-center mb-8">
      <div class="text-h3 font-weight-bold mb-2">🎲 {{ t('app.title') }}</div>
      <div class="text-body-1 text-medium-emphasis">{{ t('home.subtitle') }}</div>
    </div>

    <v-row>
      <v-col v-for="item in items" :key="item.to" cols="12" sm="6" md="3">
        <v-card
          :to="item.to"
          rounded="xl"
          hover
          class="text-center pa-6"
          elevation="1"
        >
          <div class="text-h2 mb-3">{{ item.emoji }}</div>
          <div class="text-h6 font-weight-bold">{{ t(item.label) }}</div>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const { t } = useI18n()

const items = [
  { to: '/coin', emoji: '🪙', label: 'nav.coin' },
  { to: '/dice', emoji: '🎲', label: 'nav.dice' },
  { to: '/card', emoji: '🃏', label: 'nav.card' },
  { to: '/number', emoji: '🎯', label: 'nav.number' },
]
</script>
