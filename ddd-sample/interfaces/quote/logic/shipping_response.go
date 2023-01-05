/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    response
	@Date    2022/5/30 11:29
	@Desc
*/

package logic

import (
	"sort"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	qtLogic "github.com/heshaofeng1991/ddd-sample/interfaces/quote"
)

func BuildQuoteResponse(shippingFees []*domainEntity.ShippingFeeForFilterCost) *qtLogic.QuoteRsp {
	quoteRsp := &qtLogic.QuoteRsp{
		Data: &qtLogic.QuoteInfo{},
	}

	if len(shippingFees) < 3 {
		for _, fee := range shippingFees {
			quoteRsp.Data.List = append(quoteRsp.Data.List, coverFee(fee))
		}

		return quoteRsp
	}

	cheap := getTheCostCheapest(shippingFees)
	if cheap != nil {
		quoteRsp.Data.List = append(quoteRsp.Data.List, coverFee(cheap))
	}

	fastest := getTheCostFastest(shippingFees)
	if fastest != nil {
		quoteRsp.Data.List = append(quoteRsp.Data.List, coverFee(fastest))
	}

	sort.SliceStable(
		shippingFees, func(i, j int) bool {
			return shippingFees[i].Fee().TotalFee() < shippingFees[j].Fee().TotalFee()
		})

	for _, fee := range shippingFees {
		if fee.ChannelID() == cheap.ChannelID() ||
			fee.ChannelID() == fastest.ChannelID() {
			continue
		}

		quoteRsp.Data.List = append(quoteRsp.Data.List, coverFee(fee))
	}

	return quoteRsp
}

func coverFee(fee *domainEntity.ShippingFeeForFilterCost) qtLogic.Quote {
	return qtLogic.Quote{
		ActualWeight:    fee.ActualWeight(),
		ChannelCostId:   fee.ChannelCostID(),
		ChannelId:       fee.ChannelID(),
		ChannelName:     fee.ChannelName(),
		Currency:        fee.Currency(),
		FuelFee:         fee.Fee().FuelFee(),
		MaxNormalDays:   fee.MaxNormalDays(),
		MinNormalDays:   fee.MinNormalDays(),
		MiscFee:         fee.Fee().MiscFee(),
		ProcessingFee:   fee.Fee().ProcessingFee(),
		RegistrationFee: fee.Fee().RegistrationFee(),
		ShippingFee:     fee.Fee().ShippingFee(),
		TotalFee:        fee.Fee().TotalFee(),
		LogisticsType:   fee.ChannelType(),
		VolumeWeight:    fee.VolumeWeight(),
		ChargeWeight:    fee.ChargeWeight(),
		IsRecommended:   fee.IsRecommended(),
		ChannelType:     fee.ChannelType(),
		Description:     fee.Description(),
		Tag:             fee.Tag(),
	}
}

// 找到最低费用.
func getTheCostCheapest(
	shippingFeeForFilterCosts []*domainEntity.ShippingFeeForFilterCost,
) *domainEntity.ShippingFeeForFilterCost {
	if len(shippingFeeForFilterCosts) == 0 {
		return nil
	}

	sort.SliceStable(
		shippingFeeForFilterCosts, func(i, j int) bool {
			return shippingFeeForFilterCosts[i].Fee().TotalFee() < shippingFeeForFilterCosts[j].Fee().TotalFee()
		})

	return shippingFeeForFilterCosts[0]
}

// 找到时效最快费用.
func getTheCostFastest(
	shippingFeeForFilterCosts []*domainEntity.ShippingFeeForFilterCost,
) *domainEntity.ShippingFeeForFilterCost {
	if len(shippingFeeForFilterCosts) == 0 {
		return nil
	}

	sort.SliceStable(
		shippingFeeForFilterCosts, func(i, j int) bool {
			return shippingFeeForFilterCosts[i].AverageDays() < shippingFeeForFilterCosts[j].AverageDays()
		})

	return shippingFeeForFilterCosts[0]
}
