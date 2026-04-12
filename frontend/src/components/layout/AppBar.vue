<template>
  <v-app-bar elevation="0" border="b">
    <v-app-bar-nav-icon @click="$emit('toggle-drawer')" class="d-md-none" />

    <v-app-bar-title>
      <router-link to="/" class="text-decoration-none text-on-surface font-weight-bold">
        🎲 {{ t('app.title') }}
      </router-link>
    </v-app-bar-title>

    <template #append>
      <!-- Language switcher -->
      <v-btn-toggle
        v-model="currentLocale"
        density="compact"
        rounded="pill"
        class="mr-2"
      >
        <v-btn
          v-for="loc in locales"
          :key="loc.code"
          :value="loc.code"
          size="small"
          @click="setLocale(loc.code)"
        >
          {{ loc.label }}
        </v-btn>
      </v-btn-toggle>

      <!-- Theme toggle -->
      <v-btn :icon="isDark ? 'mdi-weather-sunny' : 'mdi-weather-night'" @click="toggle" />
    </template>
  </v-app-bar>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { useI18n } from 'vue-i18n'
import { useTheme } from '@/composables/useTheme'
import { useLocale } from '@/composables/useLocale'

defineEmits<{ 'toggle-drawer': [] }>()

const { t } = useI18n()
const { isDark, toggle } = useTheme()
const { locale, locales, setLocale } = useLocale()

const currentLocale = computed(() => locale.value)
</script>
