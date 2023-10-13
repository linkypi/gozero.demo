package utils

import (
	"github.com/golang-jwt/jwt/v4"
	"myapi/internal/config"
	"myapi/internal/types"
	"time"
)

type JwtUtil struct {
	Config config.Config
}

func GenJwtToken(payloads map[string]interface{}, config config.Config) (*types.JwtTokenResponse, error) {

	secret := config.Auth.AccessSecret
	expires := config.Auth.AccessExpire
	claims := make(jwt.MapClaims)
	now := time.Now().Unix()
	claims["exp"] = now + expires
	claims["iat"] = now
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	accessToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}
	return &types.JwtTokenResponse{
		AccessToken:  accessToken,
		AccessExpire: now + expires,
		RefreshAfter: now + expires/2,
	}, nil
}
