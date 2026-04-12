import { ref } from 'vue'
import { defineStore } from 'pinia'
import { rollDice, rollMultipleDice } from '@/api/dice'
import type { DiceRollResult, MultiDiceRollResult } from '@/api/types'

export const useDiceStore = defineStore('dice', () => {
  const lastRoll = ref<DiceRollResult | null>(null)
  const lastMulti = ref<MultiDiceRollResult | null>(null)
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function roll(sides = 6) {
    loading.value = true
    error.value = null
    try {
      lastRoll.value = await rollDice(sides)
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  async function rollMultiple(count: number, sides = 6) {
    loading.value = true
    error.value = null
    try {
      lastMulti.value = await rollMultipleDice(count, sides)
    } catch (e) {
      error.value = String(e)
    } finally {
      loading.value = false
    }
  }

  return { lastRoll, lastMulti, loading, error, roll, rollMultiple }
})
