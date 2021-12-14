package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"time"
)

// 过期时间
const ONE_MONTH_SECONDS_I = 60 * 60 * 24 * 30

// Claims 自定义保存数据
type Claims struct {
	UserID    string `json:"user_id"`
	Timestamp int64  `json:"timestamp"`
	jwt.StandardClaims
}

// jwt 生成方式
func GenerateJWT(jwtSecret string, userId string) (string, error) {
	unix := time.Now().Unix()
	claims := Claims{
		UserID:         userId,
		Timestamp:      unix,
		StandardClaims: jwt.StandardClaims{ExpiresAt: unix + ONE_MONTH_SECONDS_I},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	ss, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}
	return ss, nil
}

// 解析 token
func ParseJWT(jwtSecret string, token string) error {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("jwt验证token失败")
		}
		return []byte(jwtSecret), nil
	})
	if tokenClaims != nil {
		claims, ok := tokenClaims.Claims.(*Claims)
		if !ok && !tokenClaims.Valid {
			return errors.New("非法请求")
		}
		if claims.ExpiresAt < time.Now().Unix() {
			return errors.New("token过期")
		}
	}
	return err
}
