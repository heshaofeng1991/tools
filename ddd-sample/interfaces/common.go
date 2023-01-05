/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    common
	@Date    2022/5/26 10:15
	@Desc
*/

package interfaces

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/request"
	jwtAuth "github.com/heshaofeng1991/common/util/auth"
	"github.com/heshaofeng1991/common/util/env"
	"github.com/pkg/errors"
)

func GetUserID(r *http.Request) (int64, error) {
	var claims jwtAuth.WMSClaims

	keyFunc := func(token *jwt.Token) (i interface{}, e error) { return []byte(env.JwtSecret), nil }

	token, err := request.ParseFromRequest(
		r,
		request.AuthorizationHeaderExtractor,
		keyFunc,
		request.WithClaims(&claims),
	)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	userID, err := jwtAuth.Authenticate(token.Raw)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return int64(userID), nil
}
