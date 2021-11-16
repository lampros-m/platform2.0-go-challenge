package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"gwi/platform2.0-go-challenge/environment"
	gwierrors "gwi/platform2.0-go-challenge/pkg/errors"
	gwihttp "gwi/platform2.0-go-challenge/pkg/http"

	"github.com/dgrijalva/jwt-go"
)

// // Claims : Username and jwt standard claims struct.
type Claims struct {
	ID       uint32 `json:"id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// ReadJWTToken : Reads authorization header and returns the bearer token provided.
func ReadJWTToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	parts := strings.Split(auth, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("Invalid authorization request")
	}
	return parts[1], nil
}

func Authenticator(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenStr, err := ReadJWTToken(r)
		if err != nil || tokenStr == "" {
			gwihttp.ResponseWithJSON(http.StatusUnauthorized, gwierrors.ForbiddenNew("error token must be applied"), w)
			return
		}

		config := environment.LoadConfig()
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(token *jwt.Token) (interface{}, error) {
				return config.JwtKey, nil
			})
		if err != nil {
			gwihttp.ResponseWithJSON(http.StatusInternalServerError, gwierrors.InternalServerNew("error on parsing token"), w)
			return
		}

		if !tkn.Valid {
			gwihttp.ResponseWithJSON(http.StatusUnauthorized, gwierrors.UnauthorizedNew("error token is invalid"), w)
			return
		}

		userID := claims.ID
		gwiUserIndicator := config.GwiUser
		r.Header.Add(gwiUserIndicator, strconv.Itoa(int(userID)))

		next(w, r)
	}
}
