package middleware

import (
	"net/http"
	"strings"

	"github.com/evermos/boilerplate-go/configs"
	"github.com/evermos/boilerplate-go/transport/http/response"
)

type ApiKeyAuthentication struct {
	config *configs.Config
}

const (
	HeaderApiKey = "X-Api-Key"
)

func ProvideApiKeyAuthentication(config *configs.Config) *ApiKeyAuthentication {
	return &ApiKeyAuthentication{
		config: config,
	}
}

func (a *ApiKeyAuthentication) ApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientKey := r.Header.Get(HeaderApiKey)
		if !strings.EqualFold(clientKey, a.config.App.APIKEY) {
			response.WithMessage(w, http.StatusUnauthorized, "invalid api key")
			return
		}
		w.Header().Set("X-Api-Key", clientKey)
		next.ServeHTTP(w, r)
	})
}

func (a *ApiKeyAuthentication) CustomMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		println("Custom logger middleware at", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
