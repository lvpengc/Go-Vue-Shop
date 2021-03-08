// Package jwt 对jwt-token的创建解析操作
package jwt

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/hellozouyou/easy-mall/pkg/constant"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token jwt生成的Token，包含当前用户uid在cache里的key即uuid
type Token struct {
	Exp  int64  `json:"exp"`  // 具体过期时间戳
	UUID string `json:"uuid"` // 用于从cache中获取实际用户ID
}

// TokenRes 返回给前端的token声明格式
type TokenRes struct {
	AccessToken string `json:"access_token"` // 上述Token加密生成的token
	TokenType   string `json:"token_type"`   // token的类型，由前端作为access_token的键名随Header提交
	ExpiresIn   int64  `json:"expires_in"`   // token的过期时间
}

var (
	// jwtKey 生成jwt-token的盐值字符串，需唯一以确保多进程下解析成功
	jwtKey string
)

// Init jwt 初始化
func Init() {
	jwtKey = "9v02jij+f0wreif0=wefjw01-rg43g37"
}

// Create 根据Token加盐生成token字符串
func Create(m *Token) (authToken TokenRes, err error) {
	// 初始化jwt
	token := jwt.New(jwt.SigningMethodHS256)
	// 数据转json解析到jwt
	str, err := json.Marshal(m)
	if err != nil {
		return
	}
	err = json.Unmarshal(str, &token.Claims)
	if err != nil {
		return
	}
	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return
	}
	authToken.AccessToken = tokenString
	// 2小时有效期
	authToken.ExpiresIn = time.Now().Add(time.Hour * time.Duration(2)).Unix()
	authToken.TokenType = constant.TOKENKEY
	return
}

// Parse 解析token
func Parse(tokenString string) (Token, error) {
	var (
		t   Token
		err error
	)
	// 将token字符串解析
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jwtKey), nil
	})
	if err != nil {
		return t, err
	}
	// 判断token是否有效
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// 将token中数据解析
		val, err := json.Marshal(claims)
		if err != nil {
			return t, err
		}
		t = Token{}

		if err = json.Unmarshal(val, &t); err != nil {
			return t, err
		}
		return t, nil
	}
	return t, errors.New("invalid token")
}
