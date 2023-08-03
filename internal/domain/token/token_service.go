package token

import "github.com/evermos/boilerplate-go/shared/jwt"

type TokenServiceImpl struct {
}

type JwtPayload struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

type JwtResponseFormat struct {
	AccessToken string `json:"access_token"`
}

type TokenService interface {
	Generate(payload JwtPayload) (JwtResponseFormat, error)
	Validate(tokenString string) error
}

func NewTokenServiceImpl() *TokenServiceImpl {
	return &TokenServiceImpl{}
}

func (s *TokenServiceImpl) Generate(payload JwtPayload) (JwtResponseFormat, error) {
	gen := jwt.NewJWT()
	tokenString, err := gen.GenerateJwt(payload.UserId, payload.UserName)
	if err != nil {
		return JwtResponseFormat{}, err
	}
	return JwtResponseFormat{AccessToken: tokenString}, nil
}

func (s *TokenServiceImpl) Validate(tokenString string) error {
	val := jwt.NewJWT()
	_, err := val.ValidateJwt(tokenString)
	if err != nil {
		return err
	}
	return nil
}
