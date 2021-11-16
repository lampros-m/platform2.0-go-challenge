package authentication

import (
	"context"
)

// Repository : Interface for authentication repository.
type Repository interface {
	Signup(ctx context.Context, credentials Credentials) error
	GetUser(ctx context.Context, username string) (User, error)
}
