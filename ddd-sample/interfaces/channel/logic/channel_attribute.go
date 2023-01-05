/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_attribute
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
	"github.com/sirupsen/logrus"
)

func (h HTTPServer) GetAttributes(w http.ResponseWriter, r *http.Request) {
	rsp := chlLogic.QueryAttributeRsp{
		Data: &chlLogic.QueryAttributeInfo{},
	}

	results, err := h.application.Queries.GetChannelAttributes.Handle(context.Background())
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	for _, val := range results {
		rsp.Data.List = append(rsp.Data.List, chlLogic.QueryAttributeData{
			Attribute: val.Attribute(),
		})
	}

	render.Respond(w, r, rsp)

	return
}

func (h HTTPServer) PostAttributes(w http.ResponseWriter, r *http.Request) {
	rsp := chlLogic.CreateAttributeRsp{
		Data: &chlLogic.CreateAttributeData{},
	}

	body := chlLogic.CreateAttributeRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logrus.Errorf("Error %v, body %v", err, r.Body)

		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	result, err := h.application.Commands.CreateChannelAttribute.Handle(context.Background(), *body.Attribute)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	rsp.Data.Status = &result

	render.Respond(w, r, rsp)

	return
}

func (h HTTPServer) PutAttributes(w http.ResponseWriter, r *http.Request) {
	rsp := chlLogic.UpdateAttributeRsp{
		Data: &chlLogic.UpdateAttributeData{},
	}

	body := chlLogic.BatchUpdateRequestBody{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		httpresponse.ErrRsp(w, r, http.StatusBadRequest, "invalid request params"+err.Error())

		return
	}

	results, err := h.application.Commands.
		UpdateChannelAttribute.
		Handle(context.Background(), *body.Attribute, *body.ChannelIds)
	if err != nil {
		httpresponse.ErrRsp(w, r, http.StatusInternalServerError, "internal server error"+err.Error())

		return
	}

	list := make([]chlLogic.UpdateAttributeInfo, 0)

	for _, val := range results {
		channelID := val.ChannelID()
		reason := val.Reason()
		status := val.Status()

		list = append(list, chlLogic.UpdateAttributeInfo{
			ChannelId: &channelID,
			Reason:    &reason,
			Status:    &status,
		})
	}

	rsp.Data.List = &list

	render.Respond(w, r, rsp)

	return
}
