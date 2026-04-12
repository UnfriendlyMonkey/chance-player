import { ref } from 'vue'
import { defineStore } from 'pinia'
import { drawCard } from '@/api/card'
import type { CardDrawResult } from '@/api/types'

export const useCardStore = defineStore('card', () => {
  const lastDraw = ref<CardDrawResult | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function draw() {
    loading.value = true
    error.value = null
    try {
      lastDraw.value = await drawCard()
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  return { lastDraw, loading, error, draw }
})
