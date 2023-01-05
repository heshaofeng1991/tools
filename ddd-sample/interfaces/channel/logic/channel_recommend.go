/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend
	@Date    2022/5/25 11:05
	@Desc
*/

package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/heshaofeng1991/common/util/httpresponse"
	chlLogic "github.com/heshaofeng1991/ddd-sample/interfaces/channel"
)

func (h HTTPServer) GetRecommends(w http.ResponseWriter, r *http.Request, params chlLogic.GetRecommendsParams) {
	pageSize, current := getPageSizeAndCurrent(params.Current, params.PageSize)

	results, total, err := h.application.Queries.GetChannelRecommends.Handle(
		context.Background(),
		params.CountryCode,
		params.Sorter,
		params.ChannelId,
		(*bool)(params.Status),
		(*bool)(params.IsRecommended),
		current,
		pageSize,
	)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error query channel recommends"+err.Error())

		return
	}

	rsp := chlLogic.QueryRecommendRsp{
		Data: &chlLogic.QueryRecommendInfo{
			Meta: chlLogic.MetaData{},
		},
	}

	for _, val := range results {
		rsp.Data.List = append(rsp.Data.List, chlLogic.QueryRecommendData{
			ChannelId:     val.ChannelID(),
			CountryCode:   val.CountryCode(),
			IsRecommended: val.IsRecommended() == 1,
			Status:        val.Status() == 1,
			UpdatedAt:     val.UpdatedAt().Format(UTC),
		})
	}

	rsp.Data.Meta.Current = current
	rsp.Data.Meta.PageSize = pageSize
	rsp.Data.Meta.Total = total

	render.Respond(w, r, rsp)

	return
}

//nolint:revive
func (h HTTPServer) PostIdRecommends(w http.ResponseWriter, r *http.Request, id int64) {
	body := chlLogic.CreateRecommendRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	result, err := h.application.Commands.CreateChannelRecommend.Handle(
		context.Background(),
		body.CountryCode,
		id)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "create channel recommend"+err.Error())

		return
	}

	render.Respond(w, r, chlLogic.CreateRecommendRsp{
		Data: &chlLogic.CreateRecommendData{
			Status: int32(result),
		},
	})

	return
}

//nolint:revive
func (h HTTPServer) PutIdRecommends(w http.ResponseWriter, r *http.Request, id int64) {
	body := chlLogic.UpdateRecommendRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	result, err := h.application.Commands.UpdateChannelRecommend.Handle(
		context.Background(),
		body.CountryCode,
		id,
		bool(body.IsRecommended),
		bool(body.Status))
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "update channel recommends"+err.Error())

		return
	}

	render.Respond(w, r, chlLogic.CreateRecommendRsp{
		Data: &chlLogic.CreateRecommendData{
			Status: int32(result),
		},
	})

	return
}
