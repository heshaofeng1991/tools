/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend
	@Date    2022/5/30 11:33
	@Desc
*/

package logic

import (
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

func getChannelIDs(fees []*domainEntity.ShippingFeeForFilterCost) []int64 {
	result := make([]int64, 0)

	for _, val := range fees {
		result = append(result, val.ChannelID())
	}

	return result
}

func setChannelRecommend(fees []*domainEntity.ShippingFeeForFilterCost,
	channelID int64,
) []*domainEntity.ShippingFeeForFilterCost {
	for _, val := range fees {
		if val.ChannelID() == channelID {
			val = domainEntity.NewShippingFeeForFilterCost(
				*domainEntity.NewShippingFee(
					val.Fee().TotalFee(),
					val.Fee().FuelFee(),
					val.Fee().ProcessingFee(),
					val.Fee().RegistrationFee(),
					val.Fee().MiscFee(),
					val.Fee().ShippingFee(),
				),
				val.Mode(),
				val.ChannelCostID(),
				val.ChannelID(),
				val.ChannelName(),
				"USD",
				val.ChargeWeight(),
				val.ActualWeight(),
				val.VolumeWeight(),
				val.LogisticsType(),
				val.MinNormalDays(),
				val.MaxNormalDays(),
				true,
				val.ChannelType(),
				val.DeliverDuty(),
				*val.Description(),
				*val.Tag(),
				val.AverageDays(),
			)

			break
		}
	}

	return fees
}
