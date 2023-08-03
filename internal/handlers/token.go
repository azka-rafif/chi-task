package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/evermos/boilerplate-go/internal/domain/token"
	"github.com/evermos/boilerplate-go/shared/failure"
	"github.com/evermos/boilerplate-go/transport/http/middleware"
	"github.com/evermos/boilerplate-go/transport/http/response"
	"github.com/go-chi/chi"
)

type TokenHandler struct {
	Service token.TokenService
	JwtAuth *middleware.JwtAuthentication
}

func NewTokenHandler(service token.TokenService) TokenHandler {
	return TokenHandler{Service: service}
}

func (h *TokenHandler) Router(r chi.Router) {
	r.Route("/token", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", h.Generate)
		})
		r.Group(func(r chi.Router) {
			r.Use(h.JwtAuth.Validate)
			r.Get("/validate", h.Validate)
		})
	})
}

func (h *TokenHandler) Generate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var payload token.JwtPayload
	err := decoder.Decode(&payload)
	if err != nil {
		response.WithError(w, failure.BadRequest(err))
		return
	}
	res, err := h.Service.Generate(payload)
	if err != nil {
		response.WithError(w, err)
		return
	}
	response.WithJSON(w, http.StatusCreated, res)
}

func (h *TokenHandler) Validate(w http.ResponseWriter, r *http.Request) {
	response.WithMessage(w, http.StatusOK, "Success")
}
