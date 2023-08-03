package middleware

import (
	"net/http"

	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/shared/jwt"
	"github.com/evermos/boilerplate-go/transport/http/response"
)

type JwtAuthentication struct {
	config *configs.Config
}

const (
	HeaderJwt = "Authorization"
)

func ProvideJwtAuthentication(config *configs.Config) *JwtAuthentication {
	return &JwtAuthentication{
		config: config,
	}
}

func (a *JwtAuthentication) Validate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get(HeaderJwt)
		if token == "" {
			response.WithMessage(w, http.StatusUnauthorized, "invalid jwt")
			return
		}
		validator := jwt.NewJWT()
		claims, err := validator.ValidateJwt(token)
		if err != nil {
			response.WithMessage(w, http.StatusUnauthorized, "invalid jwt")
			return
		}
		println(claims.UserName, claims.UserId)
		next.ServeHTTP(w, r)
	})
}
