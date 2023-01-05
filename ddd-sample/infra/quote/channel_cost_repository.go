/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
	@Date    2022/5/12 10:08
	@Desc
*/

package quote

import (
	"context"
	"math"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcost"
	"github.com/heshaofeng1991/entgo/ent/gen/predicate"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetChannelCostByWeight(ctx context.Context, batchIDs []int64,
	countryCode string, actualGrams int,
) ([]*domainEntity.ChannelCost, error) {
	var (
		err                error
		weightChannelCosts []*ent.ChannelCost
	)

	weightChannelCosts, err = c.entClient.ChannelCost.Query().Where(
		channelcost.And(
			channelcost.ChannelCostBatchIDIn(batchIDs...),
			channelcost.CountryCode(countryCode),
			channelcost.StartWeightLT(actualGrams),
			channelcost.EndWeightGTE(actualGrams),
			channelcost.DeletedAtIsNil(),
		),
	).Order(ent.Desc(channelcost.FieldUpdatedAt)).WithChannels().All(ctx)

	return domainEntity.CovertChannelCost(weightChannelCosts), errors.Wrap(err, "")
}

func (c ShippingOptionRepository) QueryChannelCostByVolume(ctx context.Context, countryCode string,
	chlCosts []predicate.ChannelCost,
) ([]*ent.ChannelCost, error) {
	var (
		err                error
		volumeChannelCosts []*ent.ChannelCost
	)

	volumeChannelCosts, err = c.entClient.ChannelCost.Query().Where(channelcost.And(
		channelcost.Or(chlCosts...),
		channelcost.CountryCode(countryCode),
		channelcost.Status(1),
		channelcost.DeletedAtIsNil(),
	),
	).Order(ent.Desc(channelcost.FieldUpdatedAt)).WithChannels().All(ctx)

	return volumeChannelCosts, errors.Wrap(err, "")
}

func (c ShippingOptionRepository) GetChannelCostByVolume(ctx context.Context, countryCode string,
	actualGrams int, volume int64, volumeChannelCostBatches []*domainEntity.ChannelBatch,
) (rsp []*domainEntity.ChannelCostRsp, err error) {
	var (
		chlCost predicate.ChannelCost
		factor  float64
	)

	chlCosts := make([]predicate.ChannelCost, 0)

	mapVols := filterVolumeData(int32(volume), volumeChannelCostBatches)

	for _, val := range mapVols {
		chlCost = channelcost.And(
			channelcost.ChannelCostBatchIDIn(val.BatchID...),
			channelcost.StartWeightLT(int(val.VolumeWeight)),
			channelcost.EndWeightGTE(int(val.VolumeWeight)))

		chlCosts = append(chlCosts, chlCost)
	}

	volumeChannelCosts, err := c.QueryChannelCostByVolume(ctx, countryCode, chlCosts)
	if err != nil {
		return rsp, errors.Wrap(err, "")
	}

	for _, v := range volumeChannelCosts {
		factor = math.Ceil(float64(volume) / float64(v.Edges.Channels.VolumeFactor))

		channelCosts := make([]*ent.ChannelCost, 0)
		channelCosts = append(channelCosts, v)

		channelCostRsp := domainEntity.NewChannelCostRsp(
			domainEntity.CovertChannelCost(channelCosts)[0],
			int64(factor),
			int64(actualGrams),
			int64(factor))

		rsp = append(rsp, channelCostRsp)
	}

	return rsp, errors.Wrap(err, "")
}

func filterVolumeData(volume int32, volumeChannelCostBatches []*domainEntity.ChannelBatch) []domainEntity.MapVols {
	var (
		vol    domainEntity.Vols
		mapVol domainEntity.MapVols
		factor float64
	)

	vols := make([]domainEntity.Vols, 0)
	uniqueVols := make([]int, 0)
	mapVols := make([]domainEntity.MapVols, 0)

	for _, volumeChannelCostBatch := range volumeChannelCostBatches {
		if volumeChannelCostBatch.ChannelData().VolumeFactor() > 0 {
			factor = math.Ceil(float64(volume) / float64(volumeChannelCostBatch.ChannelData().VolumeFactor()))
			vol = domainEntity.Vols{
				VolumeWeight: int64(factor),
				BatchID:      volumeChannelCostBatch.ChannelCostBatch().ID(),
			}

			vols = append(vols, vol)
		}
	}

	mapped := make(map[int]bool)

	for _, value := range vols {
		if _, ok := mapped[int(value.VolumeWeight)]; !ok {
			mapped[int(value.VolumeWeight)] = true

			uniqueVols = append(uniqueVols, int(value.VolumeWeight))
		}
	}

	for _, uniqueVol := range uniqueVols {
		mapVol = domainEntity.MapVols{
			VolumeWeight: int64(uniqueVol),
		}

		for _, data := range vols {
			if data.VolumeWeight == int64(uniqueVol) {
				mapVol.BatchID = append(mapVol.BatchID, data.BatchID)
			}
		}

		mapVols = append(mapVols, mapVol)
	}

	return mapVols
}
