// Package user defines the User domain entity and its repository interface.
// The repository is not yet implemented — this is a forward declaration
// so the rest of the codebase can reference it when auth is added.
package user

import "context"

// User represents an authenticated platform user.
type User struct {
	ID         string
	Name       string
	TelegramID int64 // set when account linked to Telegram
}

// Repository is the persistence interface for users.
// Implement in infra/repository when a database is introduced.
type Repository interface {
	FindByID(ctx context.Context, id string) (*User, error)
	Save(ctx context.Context, u *User) error
}
