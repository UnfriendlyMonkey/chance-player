import { get } from './client'
import type { RandomNumberResult } from './types'

export const randomNumber = (min = 1, max = 100) =>
  get<RandomNumberResult>('/number/random', { min: String(min), max: String(max) })
