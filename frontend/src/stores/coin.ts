import { ref } from 'vue'
import { defineStore } from 'pinia'
import { flipCoin, flipCoins, consultIChing } from '@/api/coin'
import type { CoinFlipResult, MultiCoinFlipResult, IChingConsultResult } from '@/api/types'

export const useCoinStore = defineStore('coin', () => {
  const lastFlip = ref<CoinFlipResult | null>(null)
  const lastMulti = ref<MultiCoinFlipResult | null>(null)
  const iching = ref<IChingConsultResult | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function flip() {
    loading.value = true
    error.value = null
    try {
      lastFlip.value = await flipCoin()
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  async function flipMultiple(count: number) {
    loading.value = true
    error.value = null
    try {
      lastMulti.value = await flipCoins(count)
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  async function iChing() {
    loading.value = true
    error.value = null
    try {
      iching.value = await consultIChing()
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  return { lastFlip, lastMulti, iching, loading, error, flip, flipMultiple, iChing }
})
