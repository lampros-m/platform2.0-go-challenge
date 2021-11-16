package authentication

import (
	"context"
	"errors"
	"fmt"
	"time"

	"gwi/platform2.0-go-challenge/api/middleware"
	"gwi/platform2.0-go-challenge/environment"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

const (
	tokenExpirationInMinutes = 240
)

// Service : Struct that represents authentication service.
type Service struct {
	repo Repository
}

// NewService : Service dashboard constructor.
func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Signup : Signup a new user.
func (o *Service) Signup(ctx context.Context, username string, password string) error {
	var err error
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return err
	}

	credentials := Credentials{}
	credentials.Username = username
	credentials.Password = string(hashedPassword)

	err = o.repo.Signup(ctx, credentials)
	if err != nil {
		return fmt.Errorf("could not signup a user with repository error: %s", err.Error())
	}

	return nil
}

// Login : Login for a user.
func (o *Service) Login(ctx context.Context, username string, password string) (string, error) {
	var err error

	user, err := o.GetUser(ctx, username)
	if err != nil {
		return "", fmt.Errorf("could not fetch user from db with repository error: %s", err.Error())
	}

	hashedPassword := user.HashedPassword
	unHashedPassword := password

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(unHashedPassword))
	if err != nil {
		return "", fmt.Errorf("user: %s could not be identified", username)
	}

	expirationTime := time.Now().Add(tokenExpirationInMinutes * time.Minute)
	claims := &middleware.Claims{
		ID:       user.ID,
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	config := environment.LoadConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.JwtKey)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return tokenString, nil
}

// GetUser : Returns a user based on his username.
func (o *Service) GetUser(ctx context.Context, username string) (User, error) {
	return o.repo.GetUser(ctx, username)
}
