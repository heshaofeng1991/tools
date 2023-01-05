/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_config
	@Date    2022/5/30 11:33
	@Desc
*/

package logic

import (
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

func filterShippingFees(fees []*domainEntity.ShippingFeeForFilterCost,
	channelConfigs []*domainEntity.ChannelConfig,
) []*domainEntity.ShippingFeeForFilterCost {
	result := make([]*domainEntity.ShippingFeeForFilterCost, 0)

	for _, val := range fees {
		for _, cfg := range channelConfigs {
			if cfg.ChannelID() == val.ChannelID() {
				continue
			}
		}

		result = append(result, val)
	}

	return result
}
