/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_coverty_currency
	@Date    2022/5/30 11:41
	@Desc
*/

package logic

import (
	"math"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

// 进行最终汇率的转换（默认的是人民币汇率）RMB To USD.
func covertToUSDCurrency(
	shippingFeeForFilterCosts []*domainEntity.ShippingFeeForFilterCost,
) (rsp []*domainEntity.ShippingFeeForFilterCost) {
	for _, shippingFeeForFilterCost := range shippingFeeForFilterCosts {
		shippingFee := math.Ceil((shippingFeeForFilterCost.Fee().ShippingFee()/DefaultExchangeRate)*
			FeePrecision) / FeePrecision
		fuelFee := math.Ceil(shippingFeeForFilterCost.Fee().FuelFee()/DefaultExchangeRate*
			FeePrecision) / FeePrecision
		processingFee := math.Ceil(shippingFeeForFilterCost.Fee().ProcessingFee()/DefaultExchangeRate*
			FeePrecision) / FeePrecision
		registrationFee := math.Ceil(shippingFeeForFilterCost.Fee().RegistrationFee()/DefaultExchangeRate*
			FeePrecision) / FeePrecision
		miscFee := math.Ceil(shippingFeeForFilterCost.Fee().MiscFee()/DefaultExchangeRate*
			FeePrecision) / FeePrecision
		totalFee := math.Ceil((shippingFee+fuelFee+processingFee+registrationFee+miscFee)*
			FeePrecision) / FeePrecision

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
			"USD",
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
