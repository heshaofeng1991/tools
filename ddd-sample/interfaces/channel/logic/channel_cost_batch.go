/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
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

//nolint:revive
func (h HTTPServer) GetChannelIdCostBatches(w http.ResponseWriter, r *http.Request,
	channelId int64, params chlLogic.GetChannelIdCostBatchesParams,
) {
	pageSize, current := getPageSizeAndCurrent(params.Current, params.PageSize)

	rsp := chlLogic.ChannelCostBatchSearchRsp{
		Data: &chlLogic.ChannelCostBatchSearchInfo{
			Meta: chlLogic.MetaData{},
		},
	}

	var (
		batchStatus *bool
		status      bool
	)

	if params.Status != nil && *params.Status == 0 {
		status = false
		batchStatus = &status
	}

	if params.Status != nil && *params.Status == 1 {
		status = true
		batchStatus = &status
	}

	results, total, err := h.application.Queries.GetChannelCostBatches.Handle(
		context.Background(),
		params.EffectiveDate,
		params.Sorter,
		channelId,
		batchStatus,
		current,
		pageSize)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError,
			"internal error get channel cost batch"+err.Error())

		return
	}

	for _, val := range results {
		status := 0
		if val.Status() {
			status = 1
		}

		batch := chlLogic.ChannelCostBatchInfo{
			ChannelId:     val.ChannelID(),
			EffectiveDate: val.EffectiveDate().Format(UTC),
			Id:            val.ID(),
			Status:        status,
			UpdatedAt:     val.UpdatedAt().Format(UTC),
		}

		rsp.Data.List = append(rsp.Data.List, batch)
	}

	rsp.Data.Meta.Total = total
	rsp.Data.Meta.Current = current
	rsp.Data.Meta.PageSize = pageSize

	render.Respond(w, r, rsp)

	return
}

//nolint:revive
func (h HTTPServer) PostChannelIdCostBatches(w http.ResponseWriter, r *http.Request, channelId int64) {
	rsp := chlLogic.CreateChannelCostBatchRsp{
		Data: &chlLogic.CreateChannelCostBatchInfo{},
	}

	body := chlLogic.CreateChannelCostBatchRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	flag := false
	if body.Status == 1 {
		flag = true
	}

	result, err := h.application.Commands.CreateChannelCostBatch.Handle(
		context.Background(),
		body.EffectiveDate,
		channelId,
		flag,
	)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError,
			"internal error create channel cost batch"+err.Error())

		return
	}

	rsp.Data.ChannelCostBatchId = result

	render.Respond(w, r, rsp)

	return
}

//nolint:revive
func (h HTTPServer) PutChannelIdCostBatchesId(
	w http.ResponseWriter,
	r *http.Request,
	channelId int64,
	id int64,
) {
	rsp := chlLogic.UpdateChannelCostBatchRsp{
		Data: &chlLogic.UpdateChannelCostBatchInfo{},
	}

	body := chlLogic.CreateChannelCostBatchRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	flag := false
	if body.Status == 1 {
		flag = true
	}

	result, err := h.application.Commands.UpdateChannelCostBatch.Handle(
		context.Background(),
		body.EffectiveDate,
		channelId,
		id,
		flag,
	)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError,
			"internal error update channel cost batch"+err.Error())

		return
	}

	rsp.Data.Status = result

	render.Respond(w, r, rsp)

	return
}
