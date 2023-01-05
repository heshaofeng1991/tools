/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    http
	@Date    2022/5/12 11:21
	@Desc
*/

package logic

import (
	"context"
	"net/http"

	"github.com/go-chi/render"
	"github.com/heshaofeng1991/common/util/httpresponse"
	applicationQuote "github.com/heshaofeng1991/ddd-sample/application/quote"
	inter "github.com/heshaofeng1991/ddd-sample/interfaces"
	qtLogic "github.com/heshaofeng1991/ddd-sample/interfaces/quote"
	"github.com/sirupsen/logrus"
)

type HTTPServer struct {
	application applicationQuote.Application
}

func NewHTTPServer(application applicationQuote.Application) HTTPServer {
	return HTTPServer{
		application: application,
	}
}

func (h HTTPServer) Get(w http.ResponseWriter, r *http.Request, params qtLogic.GetParams) {
	userID, err := inter.GetUserID(r)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	shippingFeeData, err := h.getShippingOptions(context.Background(), params, userID)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	ids := getChannelIDs(shippingFeeData)

	// 查询是否有最推荐渠道.
	channelRecommendData, err := h.application.Queries.
		GetChannelRecommend.Handle(context.Background(),
		ids, params.DestCountry)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	// 过滤不支持国家渠道.
	channelConfigs, err := h.application.Queries.GetChannelConfigs.Handle(context.Background(),
		ids, params.DestCountry, userID)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	// 出参处理.
	rsp := BuildQuoteResponse(filterShippingFees(
		setChannelRecommend(shippingFeeData,
			channelRecommendData.ChannelID()), channelConfigs))

	// 出参打印.
	logrus.WithFields(
		logrus.Fields{
			"response": rsp,
		},
	).Infof("Quote Request Params")

	render.Respond(w, r, rsp)

	return
}
