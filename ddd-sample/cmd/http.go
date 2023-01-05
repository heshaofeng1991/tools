/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    http
	@Date    2022/5/11 16:28
	@Desc
*/

package main

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go/request"
	"github.com/go-chi/docgen"
	"github.com/heshaofeng1991/common/middleware"
	jwtAuth "github.com/heshaofeng1991/common/util/auth"
	"github.com/heshaofeng1991/common/util/env"
	httperr "github.com/heshaofeng1991/common/util/httpresponse"
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/sirupsen/logrus"
)

func RunHTTPServer(
	healthHandler,
	shippingOptionHandler,
	channelHandler http.Handler,
) {
	RunHTTPServerOnAddr(
		":80",
		healthHandler,
		shippingOptionHandler,
		channelHandler,
	)
}

func RunHTTPServerOnAddr(
	addr string,
	healthHandler,
	shippingOptionHandler,
	channelHandler http.Handler,
) {
	apiRouter := chi.NewRouter()

	middleware.SetMiddlewares(apiRouter)
	addAuthMiddleware(apiRouter)

	apiRouter.Mount("/logistics/v1/health-check", healthHandler)
	apiRouter.Mount("/logistics/v1/shipping-options", shippingOptionHandler)
	apiRouter.Mount("/logistics/v1/channels", channelHandler)

	logrus.Info("Starting HTTP server")

	docgen.PrintRoutes(apiRouter)

	_ = http.ListenAndServe(addr, apiRouter)
}

func addAuthMiddleware(router *chi.Mux) {
	router.Use(auth)
}

func auth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(rsp http.ResponseWriter, req *http.Request) {
		ctx := req.Context()

		if !strings.Contains(req.URL.Path, "/health-check") && req.Method != "OPTIONS" {
			var claims jwtAuth.WMSClaims

			jwtSecret := env.JwtSecret

			if jwtSecret == "" {
				jwtSecret = "logistics"
			}

			token, err := request.ParseFromRequest(
				req,
				request.AuthorizationHeaderExtractor,
				func(token *jwt.Token) (i interface{}, e error) {
					return []byte(jwtSecret), nil
				},
				request.WithClaims(&claims),
			)
			if err != nil {
				httperr.BadRequest("parse jwt token failed", err, rsp, req)

				return
			}

			logrus.Infof("error %v", err)

			if !token.Valid {
				httperr.BadRequest("invalid jwt signature", nil, rsp, req)

				return
			}

			ctx = context.WithValue(ctx, domainEntity.UserID, claims.ID)

			logrus.Infof("user id %v", claims.ID)
		}

		handler.ServeHTTP(rsp, req.WithContext(ctx))
	})
}
