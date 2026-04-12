<template>
  <v-card rounded="xl" class="pa-4">
    <v-card-title class="text-h6">{{ t('number.title') }}</v-card-title>
    <v-card-text>
      <div class="d-flex gap-2 mb-4">
        <v-text-field
          v-model.number="min"
          :label="t('number.range') + ' min'"
          type="number"
          density="compact"
          variant="outlined"
          hide-details
          style="max-width: 100px"
        />
        <v-text-field
          v-model.number="max"
          :label="t('number.range') + ' max'"
          type="number"
          density="compact"
          variant="outlined"
          hide-details
          style="max-width: 100px"
        />
      </div>
      <ResultCard :show="!!store.lastResult" :anim-key="animKey">
        <div class="text-h1 font-weight-bold">{{ store.lastResult?.number }}</div>
        <div class="text-caption text-medium-emphasis">
          {{ store.lastResult?.min }} – {{ store.lastResult?.max }}
        </div>
      </ResultCard>
    </v-card-text>
    <v-card-actions>
      <v-btn
        color="primary"
        variant="flat"
        :loading="store.loading"
        prepend-icon="mdi-numeric"
        @click="doGenerate"
      >
        {{ t('number.generate') }}
      </v-btn>
    </v-card-actions>
  </v-card>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useNumberStore } from '@/stores/number'
import ResultCard from '@/components/common/ResultCard.vue'

const { t } = useI18n()
const store = useNumberStore()
const animKey = ref(0)
const min = ref(1)
const max = ref(100)

async function doGenerate() {
  animKey.value++
  await store.generate(min.value, max.value)
}
</script>
