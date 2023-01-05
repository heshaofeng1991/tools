/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_markup
	@Date    2022/5/30 11:41
	@Desc
*/

package logic

import (
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

const (
	DefaultMode          = 3
	DefaultMarkUpPercent = 0.25
	DefaultExchangeRate  = 6.2
	FeePrecision         = 100
	DefaultValue         = 1
	JDMarkUpPercent      = 0.112
)

// 处理完成了的是否需要加入标记（markup）-- 对收益成本以及利润进行处理.
func addMarkupForShippingOptions(
	shippingFeeForFilterCosts []*domainEntity.ShippingFeeForFilterCost,
) (rsp []*domainEntity.ShippingFeeForFilterCost) {
	for _, shippingFeeForFilterCost := range shippingFeeForFilterCosts {
		// 操作make_up数据库表.
		shippingFee := shippingFeeForFilterCost.Fee().ShippingFee() * (DefaultValue + DefaultMarkUpPercent)
		fuelFee := shippingFeeForFilterCost.Fee().FuelFee() * (DefaultValue + DefaultMarkUpPercent)
		processingFee := shippingFeeForFilterCost.Fee().ProcessingFee() * (DefaultValue + DefaultMarkUpPercent)
		registrationFee := shippingFeeForFilterCost.Fee().RegistrationFee() * (DefaultValue + DefaultMarkUpPercent)
		miscFee := shippingFeeForFilterCost.Fee().MiscFee() * (DefaultValue + DefaultMarkUpPercent)
		totalFee := shippingFee + fuelFee + processingFee + registrationFee + miscFee

		shippingFeeForFilterCost := domainEntity.NewShippingFeeForFilterCost(
			*domainEntity.NewShippingFee(
				totalFee,
				fuelFee,
				processingFee,
				registrationFee,
				miscFee,
				shippingFee,
			),
			shippingFeeForFilterCost.Mode(),
			shippingFeeForFilterCost.ChannelCostID(),
			shippingFeeForFilterCost.ChannelID(),
			shippingFeeForFilterCost.ChannelName(),
			"RMB",
			shippingFeeForFilterCost.ChargeWeight(),
			shippingFeeForFilterCost.ActualWeight(),
			shippingFeeForFilterCost.VolumeWeight(),
			shippingFeeForFilterCost.LogisticsType(),
			shippingFeeForFilterCost.MinNormalDays(),
			shippingFeeForFilterCost.MaxNormalDays(),
			false,
			shippingFeeForFilterCost.ChannelType(),
			shippingFeeForFilterCost.DeliverDuty(),
			*shippingFeeForFilterCost.Description(),
			*shippingFeeForFilterCost.Tag(),
			shippingFeeForFilterCost.AverageDays(),
		)

		rsp = append(rsp, shippingFeeForFilterCost)
	}

	return rsp
}
