// Package shuffle defines the core abstractions for anything randomizable.
// It has zero external dependencies so the domain stays pure and testable.
package shuffle

// RandomSource abstracts the source of randomness.
// Implementations live in infra/random; tests may use a deterministic fake.
type RandomSource interface {
	// Intn returns a non-negative random int in [0, n).
	Intn(n int) int
}
