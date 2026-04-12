// Package random provides a cryptographically secure implementation of shuffle.RandomSource.
package random

import (
	"crypto/rand"
	"math/big"
)

// CryptoSource implements shuffle.RandomSource using crypto/rand.
// It is safe for concurrent use.
type CryptoSource struct{}

// Intn returns a cryptographically random non-negative integer in [0, n).
// Panics if n <= 0.
func (CryptoSource) Intn(n int) int {
	max := big.NewInt(int64(n))
	val, err := rand.Int(rand.Reader, max)
	if err != nil {
		// crypto/rand failure is unrecoverable in a game context
		panic("crypto/rand failed: " + err.Error())
	}
	return int(val.Int64())
}
