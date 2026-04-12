import { ref } from 'vue'
import { defineStore } from 'pinia'
import { randomNumber } from '@/api/number'
import type { RandomNumberResult } from '@/api/types'

export const useNumberStore = defineStore('number', () => {
  const lastResult = ref<RandomNumberResult | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function generate(min = 1, max = 100) {
    loading.value = true
    error.value = null
    try {
      lastResult.value = await randomNumber(min, max)
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  return { lastResult, loading, error, generate }
})
