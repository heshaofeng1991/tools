/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/25 11:04
	@Desc
*/

package logic

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/render"
	"github.com/heshaofeng1991/common/util/httpresponse"
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	chlLogic "github.com/heshaofeng1991/ddd-sample/interfaces/channel"
	"github.com/sirupsen/logrus"
)

func (h HTTPServer) Get(w http.ResponseWriter, r *http.Request, params chlLogic.GetParams) {
	pageSize, current := getPageSizeAndCurrent(params.Current, params.PageSize)

	results, total, err := h.application.Queries.GetChannels.Handle(
		context.Background(),
		params.ChannelPlatform,
		params.ChannelCode,
		params.ChannelCode,
		params.Sorter,
		current,
		pageSize)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	rsp := chlLogic.ChannelSearchRsp{
		Code:    0,
		Message: "",
		Data: &chlLogic.ChannelSearchInfo{
			Meta: chlLogic.MetaData{},
		},
	}

	for _, val := range results {
		maxLength := val.MaxLength()
		minLength := val.MinLength()
		maxThreeSideSum := val.MaxThreeSideSum()
		volumeFactor := int(val.VolumeFactor())

		chl := chlLogic.GetChannelInfo{
			Code:              val.Code(),
			CourierPlatform:   val.CourierPlatform(),
			DisplayName:       val.DisplayName(),
			EnName:            val.EnName(),
			HasChannelCost:    false,
			Id:                val.ID(),
			MaxLength:         &maxLength,
			MaxThreeSideSum:   &maxThreeSideSum,
			MinLength:         &minLength,
			Name:              val.Name(),
			PrepayTariff:      chlLogic.GetChannelInfoPrepayTariff(val.PrepayTariff()),
			QuotationCurrency: val.QuotationCurrency(),
			SortingPort:       val.SortingPort(),
			Status:            val.Status(),
			Test:              chlLogic.GetChannelInfoTest(val.Test()),
			Type:              chlLogic.GetChannelInfoType(val.ChannelLogisticType()),
			UpdatedAt:         val.UpdatedAt().Format(UTC),
			VolumeFactor:      &volumeFactor,
			WarehouseId:       val.WarehouseID(),
			ChannelType:       val.ChannelType(),
			Virtual:           val.Virtual(),
			DeliverDuty:       val.DeliverDuty(),
			Special:           val.Special(),
		}

		rsp.Data.List = append(rsp.Data.List, chl)
	}

	rsp.Data.Meta.Total = total
	rsp.Data.Meta.Current = current
	rsp.Data.Meta.PageSize = pageSize

	render.Respond(w, r, rsp)

	return
}

func (h HTTPServer) Post(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	rsp := chlLogic.CreateChannelRsp{
		Data: &chlLogic.CreateChannelInfo{},
	}

	body := chlLogic.CreateRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Error %v, body %v", err, r.Body)

		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	chl := CovertChannel(body, 0)

	channelID, err := h.application.Commands.CreateChannel.Handle(ctx, *chl)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error create channel"+err.Error())

		return
	}

	rsp.Data.ChannelId = channelID

	render.Respond(w, r, rsp)

	return
}

func CovertChannel(body chlLogic.CreateRequestBody, id int64) *domainEntity.Channel {
	var (
		volumeFactor                          int32
		maxLength, minLength, maxThreeSideSum int
	)

	if body.VolumeFactor != nil {
		volumeFactor = int32(*body.VolumeFactor)
	}

	if body.MinLength != nil {
		minLength = *body.MinLength
	}

	if body.MaxLength != nil {
		maxLength = *body.MaxLength
	}

	if body.MaxThreeSideSum != nil {
		maxThreeSideSum = *body.MaxThreeSideSum
	}

	options := make([]string, 0)
	options = append(options, Options[int8(body.Type)])
	excludeAttributes, _ := json.Marshal(body.ExcludeAttributes) //nolint:errchkjson
	option, _ := json.Marshal(options)                           //nolint:errchkjson

	return domainEntity.NewChannel(id, 1, body.CourierPlatform, body.Name,
		body.Code, int8(body.Type), "RMB", volumeFactor,
		body.EnName, body.DisplayName, 1, 0,
		maxLength, minLength, maxThreeSideSum,
		"", body.SortingPort, bool(body.PrepayTariff),
		body.Status, bool(body.Test), string(excludeAttributes), string(option),
		time.Time{}, time.Time{}, body.ChannelType, body.Virtual, body.DeliverDuty, body.Special)
}

var Options = map[int8]string{
	1: "Economic",
	2: "Fastest",
	3: "Recommended",
}

var Type = map[string]int8{
	"Economic":    1,
	"Fastest":     2,
	"Recommended": 3,
}

//nolint:revive
func (h HTTPServer) PutId(w http.ResponseWriter, r *http.Request, id int64) {
	ctx := context.Background()

	rsp := chlLogic.UpdateChannelRsp{
		Data: &chlLogic.UpdateChannelInfo{},
	}

	body := chlLogic.CreateRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	chl := CovertChannel(body, id)

	status, err := h.application.Commands.UpdateChannel.Handle(ctx, *chl)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal error update channel"+err.Error())

		return
	}

	rsp.Data.Status = status

	render.Respond(w, r, rsp)

	return
}
