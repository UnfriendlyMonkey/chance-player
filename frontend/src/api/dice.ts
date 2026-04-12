import { get } from './client'
import type { DiceRollResult, MultiDiceRollResult } from './types'

export const rollDice = (sides = 6) =>
  get<DiceRollResult>('/dice/roll', { sides: String(sides) })

export const rollMultipleDice = (count: number, sides = 6) =>
  get<MultiDiceRollResult>('/dice/roll', { count: String(count), sides: String(sides) })
