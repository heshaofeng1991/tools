/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    main
	@Date    2022/5/11 16:27
	@Desc
*/

package main

import (
	"net/http"
	"strings"

	"github.com/heshaofeng1991/common/dao"
	"github.com/heshaofeng1991/common/util/env"
	"github.com/heshaofeng1991/common/util/log"
	"github.com/heshaofeng1991/common/util/sentry"
	interfaceChannel "github.com/heshaofeng1991/ddd-sample/interfaces/channel" //nolint:gci
	interfaceChannelLogic "github.com/heshaofeng1991/ddd-sample/interfaces/channel/logic"
	interfaceHealth "github.com/heshaofeng1991/ddd-sample/interfaces/health-check"
	interfaceQuote "github.com/heshaofeng1991/ddd-sample/interfaces/quote"
	interfaceQuoteLogic "github.com/heshaofeng1991/ddd-sample/interfaces/quote/logic"
	svrChannel "github.com/heshaofeng1991/ddd-sample/service/channel"
	svrQuote "github.com/heshaofeng1991/ddd-sample/service/quote"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/sirupsen/logrus"

	_ "github.com/heshaofeng1991/entgo/ent/gen/runtime"
)

func main() {
	if env.SentryDsn != "" {
		sentry.Init()

		log.InitLog()
	}

	// Initialize the db.
	entClient, err := dao.Open()
	if err != nil {
		logrus.Infof("init db error: %v", err)

		return
	}

	switch strings.ToLower(env.ServerType) {
	case "http":
		RunHTTPServer(
			HealthHandler(),
			ShippingOptionHandler(entClient),
			ChannelHandler(entClient),
		)
	case "grpc":
	default:
	}
}

func HealthHandler() http.Handler {
	return interfaceHealth.Handler(interfaceHealth.NewHTTPServer())
}

// ShippingOptionHandler 费用计算服务.
func ShippingOptionHandler(entClient *ent.Client) http.Handler {
	app := svrQuote.NewApplication(entClient)

	return interfaceQuote.Handler(interfaceQuoteLogic.NewHTTPServer(app))
}

// ChannelHandler 渠道服务.
func ChannelHandler(entClient *ent.Client) http.Handler {
	app := svrChannel.NewApplication(entClient)

	return interfaceChannel.Handler(interfaceChannelLogic.NewHTTPServer(app))
}
