package sql

import (
	"context"
	"fmt"

	"gwi/platform2.0-go-challenge/internal/app/authentication"
	"gwi/platform2.0-go-challenge/internal/repositories/tables"

	sq "github.com/Masterminds/squirrel"
)

// Authentication : Indicates Authentication repository.
type Authentication struct {
	client BasicConnectionWithTransactions
}

// NewAuthenticationRepo : Authentication repository constructor.
func NewAuthenticationRepo(client BasicConnectionWithTransactions) *Authentication {
	return &Authentication{
		client: client,
	}
}

// Signup : Signs up a user to db.
func (o *Authentication) Signup(ctx context.Context, credentials authentication.Credentials) error {
	insertBuilder := sq.
		Insert(tables.GwiUsers).
		Columns(
			"username",
			"pass",
		).
		Values(
			credentials.Username,
			credentials.Password,
		)

	tx, err := o.client.Begin()
	if err != nil {
		return err
	}

	if _, err = insertBuilder.RunWith(tx).ExecContext(ctx); err != nil {
		_ = tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		_ = tx.Rollback()
		return err
	}

	return nil
}

// GetUser : Returns a user based on his username.
func (o *Authentication) GetUser(ctx context.Context, username string) (authentication.User, error) {
	where := sq.And{
		sq.Eq{"username": username},
	}

	q := sq.
		Select(
			"id",
			"username",
			"pass",
			"created_at",
		).
		From(tables.GwiUsers).
		Where(where)

	rows, err := q.
		RunWith(o.client).
		QueryContext(ctx)
	if err != nil {
		return authentication.User{}, err
	}
	defer rows.Close()

	var found bool
	user := authentication.User{}

	for rows.Next() {
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.HashedPassword,
			&user.Created_at,
		)
		if err != nil {
			return authentication.User{}, err
		}

		found = true
	}
	if rows.Err() != nil {
		return authentication.User{}, err
	}

	if !found {
		return authentication.User{}, fmt.Errorf("no user found for user: %s", username)
	}

	return user, nil
}
