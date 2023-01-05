/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
	@Date    2022/5/30 11:30
	@Desc
*/

package logic

import (
	"context"
	"math"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/pkg/errors"
)

// 查询可用的渠道报价.
func (h HTTPServer) getAvailableCosts(ctx context.Context, channelCostBatches []*domainEntity.ChannelBatch,
	countryCode string, actualGrams, length, width, height int, vol *int64,
) (rsp []*domainEntity.ChannelCostRsp, err error) {
	var (
		volume                                             int64
		weightChannelCostBatches, volumeChannelCostBatches []*domainEntity.ChannelBatch
		batchIDs                                           []int64
		weightRsp, volumeRsp                               []*domainEntity.ChannelCostRsp
	)

	volume = int64(length * width * height)

	if vol != nil {
		volume = *vol
	}

	weightChannelCostBatches, volumeChannelCostBatches,
		batchIDs = processCalTypeData(channelCostBatches, int64(actualGrams), volume)

	if len(weightChannelCostBatches) > 0 {
		weightRsp, err = h.CovertChannelCostByActualWeight(ctx, batchIDs, countryCode,
			int64(actualGrams), int64(actualGrams), volume)

		if err != nil {
			return rsp, errors.Wrap(err, "")
		}

		rsp = append(rsp, weightRsp...)
	}

	if len(volumeChannelCostBatches) > 0 {
		volumeRsp, err = h.application.Queries.
			GetChannelCostByVolume.Handle(
			ctx,
			countryCode,
			actualGrams,
			volume,
			volumeChannelCostBatches,
		)

		if err != nil {
			return rsp, errors.Wrap(err, "")
		}

		rsp = append(rsp, volumeRsp...)
	}

	return rsp, nil
}

func processCalTypeData(channelCostBatches []*domainEntity.ChannelBatch,
	actualWeight, volumeWeight int64) ([]*domainEntity.ChannelBatch,
	[]*domainEntity.ChannelBatch, []int64,
) {
	var (
		batchIDs                                           []int64
		weightChannelCostBatches, volumeChannelCostBatches []*domainEntity.ChannelBatch
		calValue                                           int64
		factor                                             float64
	)

	for _, channelCostBatch := range channelCostBatches {
		batch := domainEntity.NewChannelBatch(
			channelCostBatch.ChannelData(),
			channelCostBatch.ChannelCostBatch(),
			0,
			actualWeight,
			volumeWeight,
		)

		factor = math.Ceil(float64(volumeWeight) / float64(batch.ChannelData().VolumeFactor()))

		calValue = int64(factor)

		if batch.ChannelData().VolumeFactor() <= 0 ||
			(batch.ChannelData().VolumeFactor() > 0 && calValue <= actualWeight) {
			batch = domainEntity.NewChannelBatch(
				channelCostBatch.ChannelData(),
				channelCostBatch.ChannelCostBatch(),
				actualWeight,
				actualWeight,
				calValue,
			)

			weightChannelCostBatches = append(weightChannelCostBatches, batch)

			batchIDs = append(batchIDs, batch.ChannelCostBatch().ID())
		} else {
			batch = domainEntity.NewChannelBatch(
				channelCostBatch.ChannelData(),
				channelCostBatch.ChannelCostBatch(),
				calValue,
				actualWeight,
				calValue,
			)

			volumeChannelCostBatches = append(volumeChannelCostBatches, batch)
		}
	}

	return weightChannelCostBatches, volumeChannelCostBatches, batchIDs
}

func (h HTTPServer) CovertChannelCostByActualWeight(ctx context.Context, batchIDs []int64,
	countryCode string, chargeWeight, actualWeight, volumeWeight int64,
) (rsp []*domainEntity.ChannelCostRsp, err error) {
	var weightChannelCosts []*domainEntity.ChannelCost

	weightChannelCosts, err = h.application.Queries.GetChannelCostByWeight.
		Handle(ctx, batchIDs, countryCode, int(chargeWeight))

	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	for _, weightChannelCost := range weightChannelCosts {
		var vol int64

		factor := weightChannelCost.Channel().VolumeFactor()

		if factor > 0 {
			vol = volumeWeight / int64(factor)
		}

		channelCostRsp := domainEntity.NewChannelCostRsp(
			weightChannelCost,
			chargeWeight,
			actualWeight,
			vol,
		)
		rsp = append(rsp, channelCostRsp)
	}

	return rsp, errors.Wrap(err, "")
}
