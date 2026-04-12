// Package telegram provides the Telegram bot transport layer.
package telegram

import (
	"fmt"
	"strings"

	"github.com/unmo/chance-player/internal/domain/chance"
	"github.com/unmo/chance-player/internal/domain/shuffle"
)

var diceEmoji = map[int]string{
	1: "⚀", 2: "⚁", 3: "⚂", 4: "⚃", 5: "⚄", 6: "⚅",
}

var suitEmoji = map[shuffle.Suit]string{
	shuffle.Hearts:   "♥",
	shuffle.Diamonds: "♦",
	shuffle.Spades:   "♠",
	shuffle.Clubs:    "♣",
}

func formatCoinFlip(r chance.CoinFlipResult) string {
	if r.Side == shuffle.Heads {
		return "🪙 Heads"
	}
	return "🪙 Tails"
}

func formatMultiCoin(r chance.MultiCoinFlipResult) string {
	parts := make([]string, len(r.Sides))
	for i, s := range r.Sides {
		if s == shuffle.Heads {
			parts[i] = "H"
		} else {
			parts[i] = "T"
		}
	}
	return "🪙 " + strings.Join(parts, " | ")
}

func formatDiceRoll(r chance.DiceRollResult) string {
	emoji := diceEmoji[r.Value]
	if emoji == "" {
		emoji = "🎲"
	}
	return fmt.Sprintf("%s %d", emoji, r.Value)
}

func formatMultiDice(r chance.MultiDiceRollResult) string {
	parts := make([]string, len(r.Rolls))
	for i, roll := range r.Rolls {
		emoji := diceEmoji[roll.Value]
		if emoji == "" {
			emoji = "🎲"
		}
		parts[i] = fmt.Sprintf("%s %d", emoji, roll.Value)
	}
	return strings.Join(parts, "  ") + fmt.Sprintf("\nSum: %d", r.Sum)
}

func formatCard(r chance.CardDrawResult) string {
	suit := suitEmoji[r.Suit]
	if suit == "" {
		suit = string(r.Suit)
	}
	return fmt.Sprintf("🃏 %s %s", suit, r.Value)
}

func formatNumber(r chance.RandomNumberResult) string {
	return fmt.Sprintf("🎯 %d  (range %d–%d)", r.Number, r.Min, r.Max)
}

func formatIChing(r chance.IChingConsultResult) string {
	var sb strings.Builder
	sb.WriteString("📖 I Ching Consultation\n\n")
	for i, line := range r.Lines {
		coins := make([]string, len(line.Coins))
		for j, c := range line.Coins {
			if c == shuffle.Heads {
				coins[j] = "H"
			} else {
				coins[j] = "T"
			}
		}
		lineChar := "—"
		if line.LineType == chance.OldYin || line.LineType == chance.YoungYin {
			lineChar = "- -"
		}
		sb.WriteString(fmt.Sprintf("Line %d: %s  [%s] (sum=%d)\n",
			i+1, lineChar, strings.Join(coins, " "), line.Sum))
	}
	return sb.String()
}
