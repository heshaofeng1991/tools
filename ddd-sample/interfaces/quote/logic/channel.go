/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/30 11:29
	@Desc
*/

package logic

import (
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

// 过滤不支持的产品属性.
func filterChannelByAttributes(channelsData []*domainEntity.Channel,
	attributes []string,
) (rsp []*domainEntity.Channel) {
	rsp = make([]*domainEntity.Channel, 0)

	if len(attributes) == 0 {
		return channelsData
	}

	if len(attributes) == 1 &&
		strings.ToLower(attributes[0]) == "special" &&
		attributes[0] != "" {
		for _, channelData := range channelsData {
			if channelData.Special() == 1 {
				rsp = append(rsp, channelData)
			}
		}

		return rsp
	}

	var flag bool

	for _, channelData := range channelsData {
		flag = true

		for _, attr := range attributes {
			if strings.Contains(strings.ToLower(channelData.ExcludeAttributes()),
				strings.ToLower(attr)) && attr != "" {
				flag = false

				break
			}
		}

		if flag {
			rsp = append(rsp, channelData)
		}
	}

	return rsp
}
