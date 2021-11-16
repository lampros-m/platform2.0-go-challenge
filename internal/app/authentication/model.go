package authentication

import "time"

// Credentials : Describes the credentials provided at signup and login.
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User : Describes a user.
type User struct {
	ID             uint32    `json:"id"`
	Username       string    `json:"username"`
	HashedPassword string    `json:"hashed_password"`
	Created_at     time.Time `json:"created_at"`
}
