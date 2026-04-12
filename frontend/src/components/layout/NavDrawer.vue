<template>
  <v-navigation-drawer v-model="modelValue" temporary>
    <v-list nav>
      <v-list-item
        v-for="item in navItems"
        :key="item.to"
        :to="item.to"
        :prepend-icon="item.icon"
        :title="t(item.label)"
        @click="$emit('update:modelValue', false)"
      />
    </v-list>
  </v-navigation-drawer>
</template>

<script setup lang="ts">
import { useI18n } from 'vue-i18n'

const modelValue = defineModel<boolean>()
defineEmits<{ 'update:modelValue': [boolean] }>()

const { t } = useI18n()

const navItems = [
  { to: '/', icon: 'mdi-home', label: 'nav.home' },
  { to: '/coin', icon: 'mdi-circle-half-full', label: 'nav.coin' },
  { to: '/dice', icon: 'mdi-dice-6', label: 'nav.dice' },
  { to: '/card', icon: 'mdi-cards', label: 'nav.card' },
  { to: '/number', icon: 'mdi-numeric', label: 'nav.number' },
]
</script>
