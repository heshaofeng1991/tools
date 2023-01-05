/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    application
	@Date    2022/5/11 17:59
	@Desc
*/

package quote

import (
	"github.com/heshaofeng1991/ddd-sample/application/quote/query"
)

type Application struct {
	Queries Queries
}

type Queries struct {
	GetChannels            query.GetChannelInfoHandler
	GetChannelCostBatches  query.GetChannelCostBatchInfoHandler
	GetChannelCostByVolume query.GetChannelCostInfoByVolumeHandler
	GetChannelCostByWeight query.GetChannelCostInfoByWeightHandler
	GetCountryZones        query.GetCountryZoneHandler
	GetTenantInfo          query.GetTenantInfoHandler
	GetVatInfo             query.GetVatInfoHandler
	GetChannelRecommend    query.GetChannelRecommendHandler
	GetChannelConfigs      query.GetChannelConfigHandler
}
