package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID       uint
	NickName string
	jwt.StandardClaims
}

type JWT struct {
	method     jwt.SigningMethod
	privateKey interface{}
	publicKey  interface{}
}

func NewJWT(method jwt.SigningMethod, privateKey string, publicKey string) (*JWT, error) {
	j := &JWT{
		method: method,
	}
	if privateKey != "" {
		err := j.SetPrivateKey(privateKey)
		if err != nil {
			return nil, err
		}
	} else if publicKey != "" {
		err := j.SetPublicKey(publicKey)
		if err != nil {
			return nil, err
		}
	}
	return j, nil
}

// CreateToken 创建一个token
func (j *JWT) CreateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(j.method, claims)
	return token.SignedString(j.privateKey)
}

// SetPublicKey 设置公钥
func (j *JWT) SetPublicKey(signingKey string) error {
	key := []byte(signingKey)
	switch j.method {
	case jwt.SigningMethodRS256:
		public, err := jwt.ParseRSAPublicKeyFromPEM([]byte(signingKey))
		if err != nil {
			return err
		}
		j.publicKey = public
	default:
		j.publicKey = key
	}
	return nil
}

// SetPrivateKey 设置私钥
func (j *JWT) SetPrivateKey(signingKey string) error {
	key := []byte(signingKey)
	switch j.method {
	case jwt.SigningMethodRS256:
		private, err := jwt.ParseRSAPrivateKeyFromPEM(key)
		if err != nil {
			return err
		}
		j.privateKey = private
		j.publicKey = private.PublicKey
	default:
		j.privateKey = key
		j.publicKey = key
	}
	return nil
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string, customClaims jwt.Claims) (jwt.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, customClaims, func(token *jwt.Token) (i interface{}, e error) {
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if token != nil {
		if token.Valid {
			return token.Claims, nil
		}
	}
	return nil, jwt.ErrInvalidKey
}

// ParseTokenMap 解析token
func (j *JWT) ParseTokenMap(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.publicKey, nil
	})
	if err != nil {
		return nil, err
	}

	if token != nil {
		if token.Valid {
			return token.Claims.(jwt.MapClaims), nil
		}
	}
	return nil, jwt.ErrInvalidKey
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.privateKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(claims)
	}
	return "", jwt.ErrInvalidKey
}
