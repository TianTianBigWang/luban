/**
 * Created by zc on 2020/6/9.
 */
package ctr

import (
	"github.com/dgrijalva/jwt-go"
	"luban/pkg/errs"
)

type JwtClaims struct {
	jwt.StandardClaims
	User JwtUserInfo
}

type JwtUserInfo struct {
	UID      string `json:"uid"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// JwtCreate returns the JWT token by claims and secret
func JwtCreate(claims JwtClaims, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// JwtParse returns the claims by JWT token and secret
func JwtParse(tokenStr string, secret string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errs.New("Token: invalid signature")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JwtClaims)
	if !ok || !token.Valid {
		return nil, errs.New("Invalid Token")
	}
	return claims, nil
}
