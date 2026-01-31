package common

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// 定义一个密钥，实际开发中建议放在 config.yaml 里
var jwtKey = []byte("campus_market_secret_key_2024")

type Claims struct {
	UserID        uint   `json:"user_id"`
	WalletAddress string `json:"wallet_address"`
	jwt.RegisteredClaims
}

// ReleaseToken 生成 JWT Token
func ReleaseToken(userID uint, wallet string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 有效期 24 小时
	claims := &Claims{
		UserID:        userID,
		WalletAddress: wallet,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "campus_market",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// ParseToken 解析并验证 Token
func ParseToken(tokenString string) (*Claims, bool) {
	token, _ := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if token != nil && token.Valid {
		if claims, ok := token.Claims.(*Claims); ok {
			return claims, true
		}
	}
	return nil, false
}
