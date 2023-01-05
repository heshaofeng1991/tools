/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    application
	@Date    2022/5/12 10:17
	@Desc
*/

package channel

import (
	applicationChannel "github.com/heshaofeng1991/ddd-sample/application/channel"
	"github.com/heshaofeng1991/ddd-sample/application/channel/command"
	"github.com/heshaofeng1991/ddd-sample/application/channel/query"
	interfaceChannel "github.com/heshaofeng1991/ddd-sample/infra/channel"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
)

func NewApplication(entClient *ent.Client) applicationChannel.Application {
	return newApplication(entClient)
}

func newApplication(entClient *ent.Client) applicationChannel.Application {
	chlRepository := interfaceChannel.NewChlRepository(entClient)

	return applicationChannel.Application{
		Queries: applicationChannel.Queries{
			GetChannels:             query.NewGetChannelInfoHandler(chlRepository),
			GetChannelCostBatches:   query.NewGetChannelCostBatchInfoHandler(chlRepository),
			GetChannelCosts:         query.NewGetChannelCostInfoHandler(chlRepository),
			GetChannelCostBatchByID: query.NewGetChannelCostBatchByIDHandler(chlRepository),
			GetChannelAttributes:    query.NewGetChannelAttributeInfoHandler(chlRepository),
			GetChannelConfig:        query.NewGetChannelConfigHandler(chlRepository),
			GetChannelRecommends:    query.NewGetChannelRecommendHandler(chlRepository),
		},
		Commands: applicationChannel.Commands{
			CreateChannel:           command.NewCreateChannelHandler(chlRepository),
			UpdateChannel:           command.NewUpdateChannelHandler(chlRepository),
			CreateChannelCostBatch:  command.NewCreateChannelCostBatchHandler(chlRepository),
			UpdateChannelCostBatch:  command.NewUpdateChannelCostBatchHandler(chlRepository),
			CreateChannelCost:       command.NewCreateChannelCostHandler(chlRepository),
			CreateChannelAttribute:  command.NewCreateChannelAttributeHandler(chlRepository),
			UpdateChannelAttribute:  command.NewUpdateChannelAttributeHandler(chlRepository),
			CreateChannelConfig:     command.NewCreateChannelConfigHandler(chlRepository),
			UpdateChannelConfig:     command.NewUpdateChannelConfigHandler(chlRepository),
			CreateChannelRecommend:  command.NewCreateChannelRecommendHandler(chlRepository),
			UpdateChannelRecommend:  command.NewUpdateChannelRecommendHandler(chlRepository),
			UpdateChannelCostStatus: command.NewUpdateUpdateChannelCostStatusHandler(chlRepository),
		},
	}
}
