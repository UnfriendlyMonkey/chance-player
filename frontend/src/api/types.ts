// These types mirror the Go result structs in internal/domain/chance/results.go

export type CoinSide = 'heads' | 'tails'
export type Suit = 'hearts' | 'diamonds' | 'spades' | 'clubs'
export type CardValue = '6' | '7' | '8' | '9' | '10' | 'J' | 'Q' | 'K' | 'A'
export type IChingLineType = 'old_yin' | 'young_yang' | 'young_yin' | 'old_yang'

export interface CoinFlipResult {
  side: CoinSide
}

export interface MultiCoinFlipResult {
  sides: CoinSide[]
  count: number
}

export interface DiceRollResult {
  value: number
  sides: number
}

export interface MultiDiceRollResult {
  rolls: DiceRollResult[]
  sum: number
  sides: number
}

export interface CardDrawResult {
  suit: Suit
  value: CardValue
}

export interface RandomNumberResult {
  number: number
  min: number
  max: number
}

export interface IChingLineResult {
  coins: CoinSide[]
  sum: number
  line_type: IChingLineType
}

export interface IChingConsultResult {
  lines: IChingLineResult[]
}
