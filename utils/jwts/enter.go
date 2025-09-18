package jwts

import (
	"errors"
	"rbacAdmin/global"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type ClaimsUserInfo struct {
	UserID   uint   `json:"userID"`
	Username string `json:"username"`
	RoleList []uint `json:"roleList"`
}

type Claims struct {
	ClaimsUserInfo
	jwt.StandardClaims
}

// GetToken 生成Token
func GetToken(info ClaimsUserInfo) (string, error) {
	j := global.Config.JWT
	cla := Claims{
		ClaimsUserInfo: info,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(j.Expire)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, cla)
	return token.SignedString([]byte(j.Secret)) // 进行签名生成对应的token
}

// ParseToken 解析Token
func ParseToken(tokenString string) (*Claims, error) {
	j := global.Config.JWT
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		if claims.Issuer != j.Issuer {
			return nil, errors.New("token issuer is not valid")
		}
		return claims, nil
	}
	return nil, errors.New("token is not valid")
}
