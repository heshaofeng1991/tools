/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_config
	@Date    2022/5/25 11:07
	@Desc
*/

package logic

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/heshaofeng1991/common/util/httpresponse"
	inter "github.com/heshaofeng1991/ddd-sample/interfaces"
	chlLogic "github.com/heshaofeng1991/ddd-sample/interfaces/channel"
)

func (h HTTPServer) GetCustomerConfigs(w http.ResponseWriter, r *http.Request,
	params chlLogic.GetCustomerConfigsParams,
) {
	pageSize, current := getPageSizeAndCurrent(params.Current, params.PageSize)

	userID, err := inter.GetUserID(r)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error auth failed"+err.Error())

		return
	}

	results, total, err := h.application.Queries.GetChannelConfig.Handle(
		context.Background(),
		params.Sorter,
		params.ChannelId,
		(*bool)(params.Status),
		current,
		pageSize,
		userID,
	)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error query channel config"+err.Error())

		return
	}

	rsp := chlLogic.QueryCustomerConfigRsp{
		Data: &chlLogic.QueryCustomerConfigInfo{
			Meta: chlLogic.MetaData{},
		},
	}

	for _, val := range results {
		countryCodes := make([]string, 0)
		_ = json.Unmarshal([]byte(val.ExcludeCountryCode()), &countryCodes)

		rsp.Data.List = append(rsp.Data.List,
			chlLogic.QueryCustomerConfigData{
				ChannelId:   val.ChannelID(),
				Status:      val.Status() == 1,
				UpdatedAt:   val.UpdatedAt().Format(UTC),
				CountryCode: countryCodes,
			})
	}

	rsp.Data.Meta.Current = current
	rsp.Data.Meta.PageSize = pageSize
	rsp.Data.Meta.Total = total

	render.Respond(w, r, rsp)

	return
}

func (h HTTPServer) PostCustomerConfigs(w http.ResponseWriter, r *http.Request) {
	body := chlLogic.CreateCustomerConfigRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	userID, err := inter.GetUserID(r)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error auth failed"+err.Error())

		return
	}

	result, err := h.application.Commands.CreateChannelConfig.Handle(
		context.Background(),
		body.CountryCodes,
		body.Ids,
		userID)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "create channel config"+err.Error())

		return
	}

	render.Respond(w, r, chlLogic.CreateCustomerConfigRsp{
		Data: &chlLogic.CreateCustomerConfigData{
			Status: int32(result),
		},
	})

	return
}

func (h HTTPServer) PutCustomerConfigs(w http.ResponseWriter, r *http.Request) {
	body := chlLogic.UpdateCustomerConfigRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	userID, err := inter.GetUserID(r)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error auth failed"+err.Error())

		return
	}

	result, err := h.application.Commands.UpdateChannelConfig.Handle(
		context.Background(),
		body.CountryCodes,
		body.Ids,
		bool(body.Status),
		userID)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "update channel config"+err.Error())

		return
	}

	render.Respond(w, r, chlLogic.CreateCustomerConfigRsp{
		Data: &chlLogic.CreateCustomerConfigData{
			Status: int32(result),
		},
	})

	return
}
