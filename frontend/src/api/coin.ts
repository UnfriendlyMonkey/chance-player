import { get } from './client'
import type { CoinFlipResult, MultiCoinFlipResult, IChingConsultResult } from './types'

export const flipCoin = () => get<CoinFlipResult>('/coin/flip')

export const flipCoins = (count: number) =>
  get<MultiCoinFlipResult>('/coin/flip', { count: String(count) })

export const consultIChing = () => get<IChingConsultResult>('/coin/iching')
