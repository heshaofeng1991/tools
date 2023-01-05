/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    application
	@Date    2022/5/15 18:29
	@Desc
*/

package channel

import (
	"github.com/heshaofeng1991/ddd-sample/application/channel/command"
	"github.com/heshaofeng1991/ddd-sample/application/channel/query"
)

type Application struct {
	Queries  Queries
	Commands Commands
}

type Queries struct {
	GetChannels             query.GetChannelInfoHandler
	GetChannelCostBatches   query.GetChannelCostBatchInfoHandler
	GetChannelCosts         query.GetChannelCostInfoHandler
	GetChannelCostBatchByID query.GetChannelCostBatchByIDHandler
	GetChannelAttributes    query.GetChannelAttributeInfoHandler
	GetChannelRecommends    query.GetChannelRecommendHandler
	GetChannelConfig        query.GetChannelConfigHandler
}

type Commands struct {
	CreateChannel           command.CreateChannelHandler
	UpdateChannel           command.UpdateChannelHandler
	CreateChannelCostBatch  command.CreateChannelCostBatchHandler
	UpdateChannelCostBatch  command.UpdateChannelCostBatchHandler
	CreateChannelCost       command.CreateChannelCostHandler
	CreateChannelAttribute  command.CreateChannelAttributeHandler
	UpdateChannelAttribute  command.UpdateChannelAttributeHandler
	CreateChannelRecommend  command.CreateChannelRecommendHandler
	UpdateChannelRecommend  command.UpdateChannelRecommendHandler
	CreateChannelConfig     command.CreateChannelConfigHandler
	UpdateChannelConfig     command.UpdateChannelConfigHandler
	UpdateChannelCostStatus command.UpdateChannelCostStatusHandler
}
