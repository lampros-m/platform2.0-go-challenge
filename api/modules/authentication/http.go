package authentication

import "errors"

// SignupLoginRequest : Request struct for signup and login call.
type SignupLoginRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"pass"`
}

// IsValid : Indicates if something missing from requests signup and login.
func (o *SignupLoginRequest) IsValid() error {
	if o.Username == "" || o.Password == "" {
		return errors.New("username or/and password on signup request are empty")
	}
	return nil
}
