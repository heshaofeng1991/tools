/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cal
	@Date    2022/5/30 11:40
	@Desc
*/

package logic

import (
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/sirupsen/logrus"
)

// 计算成本运费.
func calculateShippingFeeForFilterCosts(channelCosts []*domainEntity.ChannelCostRsp) (
	rsp []*domainEntity.ShippingFeeForFilterCost,
) {
	for _, channelCost := range channelCosts {
		shippingFee := CalculateShippingFee(channelCost)

		if shippingFee != nil {
			if channelCost.ChannelCost().Mode() < 1 || channelCost.ChannelCost().Mode() > int8(DefaultMode) {
				continue
			}

			shippingFeeForFilterCost := domainEntity.NewShippingFeeForFilterCost(
				*domainEntity.NewShippingFee(
					shippingFee.TotalFee(),
					shippingFee.FuelFee(),
					shippingFee.ProcessingFee(),
					shippingFee.RegistrationFee(),
					shippingFee.MiscFee(),
					shippingFee.ShippingFee(),
				),
				channelCost.ChannelCost().Mode(),
				channelCost.ChannelCost().ID(),
				channelCost.ChannelCost().ChannelID(),
				channelCost.ChannelCost().Channel().DisplayName(),
				"RMB",
				channelCost.ChargeWeight(),
				channelCost.ActualWeight(),
				channelCost.VolumeWeight(),
				channelCost.ChannelCost().Channel().ChannelLogisticType(),
				channelCost.ChannelCost().MinNormalDays(),
				channelCost.ChannelCost().MaxNormalDays(),
				false,
				channelCost.ChannelCost().Channel().ChannelType(),
				channelCost.ChannelCost().Channel().DeliverDuty(),
				channelCost.ChannelCost().Channel().Description(),
				"",
				channelCost.ChannelCost().AverageDays(),
			)

			rsp = append(rsp, shippingFeeForFilterCost)
		}
	}

	return rsp
}

const (
	FeeTotalPrice               = 1 // 总价模式.
	FeeUnitPrice                = 2 // 单价模式(取整).
	FeeContinuedUnitPrice       = 3 // 续单价模式(取整)， 首重+续重.
	FeeTotalOrUnitPrice         = 4 // 总价或单价模式(取整).
	FeeUnitPriceNoCeil          = 5 // 单价模式(不取整).
	FeeContinuedUnitPriceNoCeil = 6 // 续单价模式 (不取整).
	FeeTotalOrUnitPriceNoCeil   = 7 // 总价或单价模式 (不取整).
)

func CalculateShippingFee(channelCost *domainEntity.ChannelCostRsp) (rsp *domainEntity.ShippingFee) {
	var (
		shippingCost, unitWeightFee, firstWeightFee float64
		chargeWeight, unitWeight, firstWeight       int
		mode                                        int8
	)

	chlCost := domainEntity.NewChannelCostRsp(
		channelCost.ChannelCost(),
		channelCost.ChargeWeight(),
		channelCost.ActualWeight(),
		channelCost.VolumeWeight(),
	)

	chargeWeight = int(chlCost.ChargeWeight())
	mode = chlCost.ChannelCost().Mode()
	unitWeightFee = chlCost.ChannelCost().UnitWeightFee()
	unitWeight = chlCost.ChannelCost().UnitWeight()
	firstWeightFee = chlCost.ChannelCost().FirstWeightFee()
	firstWeight = chlCost.ChannelCost().FirstWeight()

	switch mode {
	// 总价模式  eg: 0-1kg  15元.
	case FeeTotalPrice:
		shippingCost = unitWeightFee
	case FeeUnitPrice:
		shippingCost = UnitPrice(unitWeight, chargeWeight, unitWeightFee)
	case FeeContinuedUnitPrice:
		shippingCost = ContinuedUnitPrice(unitWeight, chargeWeight, firstWeight, unitWeightFee, firstWeightFee)
	case FeeTotalOrUnitPrice:
		shippingCost = TotalOrUnitPrice(unitWeight, chargeWeight, firstWeight, unitWeightFee, firstWeightFee)
	case FeeUnitPriceNoCeil:
		shippingCost = UnitPriceNoCeil(unitWeight, chargeWeight, unitWeightFee)
	case FeeContinuedUnitPriceNoCeil:
		shippingCost = ContinuedUnitPriceNoCeil(unitWeight, chargeWeight, firstWeight, unitWeightFee, firstWeightFee)
	case FeeTotalOrUnitPriceNoCeil:
		shippingCost = TotalOrUnitPriceNoCeil(unitWeight, chargeWeight, firstWeight, unitWeightFee, firstWeightFee)
	default:
		logrus.Infoln("unknown cost mode")
	}

	rsp = domainEntity.NewShippingFee(
		shippingCost+chlCost.ChannelCost().FuelFee()+
			chlCost.ChannelCost().ProcessingFee()+channelCost.ChannelCost().RegistrationFee()+
			chlCost.ChannelCost().MiscFee(),
		chlCost.ChannelCost().FuelFee(),
		chlCost.ChannelCost().ProcessingFee(),
		chlCost.ChannelCost().RegistrationFee(),
		chlCost.ChannelCost().MiscFee(),
		shippingCost,
	)

	return rsp
}
