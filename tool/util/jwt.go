package util

import (
	"github.com/golang-jwt/jwt/v5"
	"sync"
	"template/conf"
	"time"
)

var jwtSecret []byte
var once sync.Once

const TokenExpiresAt = time.Hour * 24
const Issuer = "cjs"

func InitJwtSecret() {
	once.Do(func() {
		jwtSecret = []byte(conf.Conf.App.JwtSecret)
	})
}

type Claims struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken 生成Token
func GenerateToken(id int64, username string) (string, error) {
	claims := Claims{
		id,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiresAt)),
			Issuer:    Issuer,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(jwtSecret)
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
