package api

import (
	"core/common"
	"core/config"
	"core/response"
	"core/utils"
	"strings"
	"sync"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var (
	instanceJwt *utils.JWT
	once        sync.Once
)

func InstanceJwt(method jwt.SigningMethod, privateKey string, publicKey string) (*utils.JWT, error) {
	var err error
	once.Do(func() {
		instanceJwt, err = utils.NewJWT(method, privateKey, publicKey)
	})
	return instanceJwt, err
}

func JWTAuth(jwtConf config.Jwt, customClaims jwt.Claims) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetBool("noAuth") {
			ctx.Next()
			return
		}
		auth := ctx.Request.Header.Get("Authorization")
		s := common.New(ctx)
		tokens := strings.Split(auth, " ")
		if auth == "" || len(tokens) != 2 || strings.ToLower(tokens[0]) != "bearer" {
			s.Failed(response.TokenErr.Msgf("Authorization为空或格式不对:%s", auth))
			ctx.Abort()
			return
		}
		s.Log.Info("token信息:", auth)
		j, err := InstanceJwt(jwt.SigningMethodRS256, "", jwtConf.PublicKey)
		if err != nil {
			s.Failed(response.TokenErr.Err(errors.Wrap(err, "设置密钥失败")))
			ctx.Abort()
			return
		}
		token := tokens[1]
		claims, err := j.ParseToken(token, customClaims)
		if err != nil {
			s.Failed(response.TokenErr.Err(errors.Wrap(err, "解析token失败")))
			ctx.Abort()
			return
		}

		ctx.Set("token", token)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
