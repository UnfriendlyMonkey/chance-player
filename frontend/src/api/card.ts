import { get } from './client'
import type { CardDrawResult } from './types'

export const drawCard = () => get<CardDrawResult>('/card/draw')
