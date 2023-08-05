package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JwtManager interface {
	GenerateUserJwt(userID int64) (string, error)
	ParseToken(jwtToken string) (payload *UserClaims, err error)
}

type jwtModule struct {
	SecretKey  string
	Expiration int64
}

type Options struct {
	SignKey    string
	Expiration int64 // in hour
}

func New(o *Options) JwtManager {
	return &jwtModule{
		SecretKey:  o.SignKey,
		Expiration: o.Expiration,
	}
}

type UserClaims struct {
	UserID int64 `json:"user_id"`
	jwt.StandardClaims
}

func (m *jwtModule) GenerateUserJwt(userID int64) (string, error) {
	expirationTime := time.Now().Add(time.Duration(m.Expiration) * time.Hour)
	claims := &UserClaims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
			Subject:   fmt.Sprintf("%d", userID),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.SecretKey))
}

func (m *jwtModule) ParseToken(jwtToken string) (payload *UserClaims, err error) {
	token, err := jwt.ParseWithClaims(jwtToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.SecretKey), nil
	})
	if err != nil || !token.Valid {
		err = errors.New("invalid jwt")
		return
	}

	payload = token.Claims.(*UserClaims)

	return
}
