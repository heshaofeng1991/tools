/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    application
	@Date    2022/5/12 10:18
	@Desc
*/

package quote

import (
	applicationQuote "github.com/heshaofeng1991/ddd-sample/application/quote"
	"github.com/heshaofeng1991/ddd-sample/application/quote/query"
	infraQuote "github.com/heshaofeng1991/ddd-sample/infra/quote"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

func NewApplication(entClient *ent.Client) applicationQuote.Application {
	return newApplication(entClient)
}

func newApplication(entClient *ent.Client) applicationQuote.Application {
	repository := infraQuote.NewShippingOptionRepository(entClient)

	return applicationQuote.Application{
		Queries: applicationQuote.Queries{
			GetChannels:            query.NewGetChannelInfoHandler(repository),
			GetChannelCostBatches:  query.NewGetChannelCostBatchInfoHandler(repository),
			GetChannelCostByVolume: query.NewGetChannelCostInfoByVolumeHandler(repository),
			GetChannelCostByWeight: query.NewGetChannelCostInfoByWeightHandler(repository),
			GetCountryZones:        query.NewGetCountryZoneHandler(repository),
			GetTenantInfo:          query.NewGetTenantInfoHandler(repository),
			GetVatInfo:             query.NewGetVatInfoHandler(repository),
			GetChannelRecommend:    query.NewGetChannelRecommendHandler(repository),
			GetChannelConfigs:      query.NewGetChannelConfigHandler(repository),
		},
	}
}
