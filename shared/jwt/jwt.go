package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	jwt.StandardClaims
}

type JWT struct {
}

func NewJWT() *JWT {
	return &JWT{}
}

func (j *JWT) GenerateJwt(userId, userName string) (string, error) {
	claims := Claims{
		UserId:   userId,
		UserName: userName,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour + 1).Unix(),
			Issuer:    "Azka",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("ayam"))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func (j *JWT) ValidateJwt(tokenString string) (*Claims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("ayam"), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("JWT not valid")
}
