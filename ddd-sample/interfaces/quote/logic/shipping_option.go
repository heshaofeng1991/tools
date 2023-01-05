/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    shipping_option
	@Date    2022/5/30 11:31
	@Desc
*/

package logic

import (
	"context"
	"encoding/json"
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	qtLogic "github.com/heshaofeng1991/ddd-sample/interfaces/quote"
	"github.com/pkg/errors"
)

func (h HTTPServer) getShippingOptions(ctx context.Context,
	params qtLogic.GetParams, userID int64,
) ([]*domainEntity.ShippingFeeForFilterCost, error) {
	var (
		channelsData                                        []*domainEntity.Channel
		err                                                 error
		presetChannelIds, testChannelIds, excludeChannelIds []int64
		prepayTariff                                        bool
	)

	tenantInfo, err := h.application.Queries.GetTenantInfo.Handle(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	if params.PresetChannelIds == nil && params.TestChannelIds == nil {
		_ = json.Unmarshal([]byte(tenantInfo.TestChannelIds()), &testChannelIds)
		_ = json.Unmarshal([]byte(tenantInfo.PresetChannelIds()), &presetChannelIds)
	}

	if params.ExcludeChannelIds != nil {
		excludeChannelIds = *params.ExcludeChannelIds
	}

	if params.PresetChannelIds != nil {
		for _, val := range *params.PresetChannelIds {
			presetChannelIds = append(presetChannelIds, val)
		}
	}

	if params.TestChannelIds != nil {
		for _, val := range *params.TestChannelIds {
			presetChannelIds = append(presetChannelIds, val)
		}
	}

	if params.PrepayTariff != nil {
		prepayTariff = *params.PrepayTariff
	}

	channelsData, err = h.application.Queries.GetChannels.
		Handle(ctx, prepayTariff,
			presetChannelIds, testChannelIds,
			excludeChannelIds, params.WarehouseId)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	productAttributes := make([]string, 0)

	if params.ProductAttributes != nil {
		productAttributes = *params.ProductAttributes
	}

	// 过滤渠道数据.
	filterChannelsData := filterChannelByAttributes(channelsData, productAttributes)

	// 获取可用的批次信息.
	channelCostBatches, err := h.getAvailableChannelCostBatch(ctx, filterChannelsData, params.Date)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	// 查询可用的渠道报价.
	channelCosts, err := h.getAvailableCosts(
		ctx,
		channelCostBatches,
		params.DestCountry,
		params.Weight,
		params.Length,
		params.Width,
		params.Height,
		params.Volume)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	// 通过zone（区域）过滤渠道报价.
	zoneChannelCosts, err := h.filterCostsByCountryCodeAndZipCode(ctx, channelCosts,
		params.DestCountry, params.DestZipCode)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	// 计算成本运费.
	shippingFeeForFilterCosts := calculateShippingFeeForFilterCosts(zoneChannelCosts)

	shippingFeeAddMarkup := addMarkupForShippingOptions(shippingFeeForFilterCosts)

	if params.SettlementCurrency == nil ||
		strings.ToUpper(string(*params.SettlementCurrency)) == "USD" {
		return covertToUSDCurrency(shippingFeeAddMarkup), nil
	}

	return shippingFeeAddMarkup, nil
}
