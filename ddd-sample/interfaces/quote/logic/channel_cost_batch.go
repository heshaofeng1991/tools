/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
	@Date    2022/5/30 11:30
	@Desc
*/

package logic

import (
	"context"
	"time"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/pkg/errors"
)

// 获取可用的批次信息.
func (h HTTPServer) getAvailableChannelCostBatch(ctx context.Context,
	filterChannelsData []*domainEntity.Channel, reqDate *string) (
	rsp []*domainEntity.ChannelBatch, err error,
) {
	var date time.Time

	ids := make([]int64, 0)

	date, err = checkDate(reqDate)

	if err != nil {
		return rsp, errors.Wrap(err, "")
	}

	for _, val := range filterChannelsData {
		ids = append(ids, val.ID())
	}

	channelCostBatch, err := h.application.Queries.GetChannelCostBatches.Handle(ctx, ids, date)
	if err != nil {
		return rsp, errors.Wrap(err, "")
	}

	rsp = BuildChannelCostBatch(channelCostBatch, filterChannelsData)

	return rsp, nil
}

const (
	DATE = "2006-01-02"
)

func checkDate(reqDate *string) (time.Time, error) {
	var (
		err  error
		date time.Time
	)

	date, err = time.Parse(DATE, time.Now().Format(DATE))

	if err != nil {
		return date, errors.Wrap(err, "")
	}

	if reqDate != nil {
		date, err = time.Parse(DATE, *reqDate)

		if err != nil {
			return date, errors.Wrap(err, "")
		}
	}

	return date, nil
}

func BuildChannelCostBatch(channelCostBatch []*domainEntity.ChannelCostBatch,
	filterChannelsData []*domainEntity.Channel,
) []*domainEntity.ChannelBatch {
	var result *domainEntity.ChannelBatch

	rsp := make([]*domainEntity.ChannelBatch, 0)

	for _, chl := range channelCostBatch {
		result = domainEntity.NewChannelBatch(&domainEntity.Channel{}, chl, 0, 0, 0)

		for _, filterChl := range filterChannelsData {
			if chl.ChannelID() == filterChl.ID() {
				result = domainEntity.NewChannelBatch(filterChl, chl, 0, 0, 0)

				break
			}
		}

		rsp = append(rsp, result)
	}

	return rsp
}
