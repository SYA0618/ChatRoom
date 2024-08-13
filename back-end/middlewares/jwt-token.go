package middlewares

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

func CreateDemoToken(account string, secret []byte) string {
	// 1. 填入資訊
	now := time.Now()
	claims := new(MyClaims)
	jwtId := claims.Account + strconv.FormatInt(now.Unix(), 10)
	claims.Id = jwtId
	claims.Account = account
	claims.Issuer = "demoProject"
	claims.IssuedAt = now.Unix()
	claims.ExpiresAt = now.Add(120 * time.Second).Unix()
	// 2. 編碼，製作token
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 3. 簽名
	tokenString, _ := tokenClaims.SignedString(secret)
	return tokenString
}
