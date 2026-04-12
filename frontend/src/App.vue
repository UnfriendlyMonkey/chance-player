<template>
  <v-app>
    <AppBar @toggle-drawer="drawer = !drawer" />
    <NavDrawer v-model="drawer" />

    <!-- Desktop side nav -->
    <v-navigation-drawer permanent class="d-none d-md-flex">
      <v-list nav class="mt-2">
        <v-list-item
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          :prepend-icon="item.icon"
          :title="t(item.label)"
          rounded="xl"
        />
      </v-list>
    </v-navigation-drawer>

    <v-main>
      <RouterView />
    </v-main>
  </v-app>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { RouterView } from 'vue-router'
import { useI18n } from 'vue-i18n'
import AppBar from '@/components/layout/AppBar.vue'
import NavDrawer from '@/components/layout/NavDrawer.vue'

const { t } = useI18n()
const drawer = ref(false)

const navItems = [
  { to: '/', icon: 'mdi-home', label: 'nav.home' },
  { to: '/coin', icon: 'mdi-circle-half-full', label: 'nav.coin' },
  { to: '/dice', icon: 'mdi-dice-6', label: 'nav.dice' },
  { to: '/card', icon: 'mdi-cards', label: 'nav.card' },
  { to: '/number', icon: 'mdi-numeric', label: 'nav.number' },
]
</script>
