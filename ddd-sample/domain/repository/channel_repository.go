/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_repository
	@Date    2022/5/15 17:55
	@Desc
*/

package repository

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

type ChannelRepository interface {
	GetChannels(ctx context.Context, channelPlatform, channelCode, channelName, sorter *string,
		current, pageSize int) ([]*domainEntity.Channel, error)
	CountChannels(ctx context.Context, channelPlatform, channelCode, channelName *string) (int64, error)
	CreateChannel(ctx context.Context, channel domainEntity.Channel) (int64, error)
	UpdateChannel(ctx context.Context, channel domainEntity.Channel) (int, error)

	GetChannelCostBatches(ctx context.Context, effectiveDate, sorter *string,
		channelID int64, status *bool,
		current, pageSize int) ([]*domainEntity.ChannelCostBatch, error)
	CountChannelCostBatches(ctx context.Context, effectiveDate *string,
		channelID int64, status *bool) (int64, error)
	CreateChannelCostBatch(ctx context.Context, effectiveDate string,
		channelID int64, status bool) (int64, error)
	UpdateChannelCostBatch(ctx context.Context, effectiveDate string,
		channelID, id int64, status bool) (int, error)
	GetChannelCostBatchByID(ctx context.Context, id int64) (*domainEntity.ChannelCostBatch, error)

	CreateChannelCost(ctx context.Context, costs []*domainEntity.ChannelCost) error
	GetChannelCosts(ctx context.Context, channelID,
		channelCostBatchID int64, countryCode *string,
		current, pageSize int) ([]*domainEntity.ChannelCost, error)
	CountChannelCosts(ctx context.Context, channelID,
		channelCostBatchID int64, countryCode *string) (int64, error)
	UpdateChannelCostStatus(ctx context.Context, ids []int64,
		countryCodes []string, status bool) (int, error)

	GetChannelAttributes(ctx context.Context) ([]*domainEntity.ChannelAttribute, error)
	CreateChannelAttributes(ctx context.Context, attribute string,
		updateFn func(ctx context.Context, attribute string) (int32, error)) (int32, error)
	UpdateChannelAttributes(ctx context.Context, attribute string,
		channelIDs []int64) ([]*domainEntity.UpdateChannelAttribute, error)

	GetChannelConfigs(ctx context.Context, sorter *string, channelID *int64,
		status *bool, current, pageSize int, userID int64) ([]*domainEntity.ChannelConfig, error)
	CountChannelConfigs(ctx context.Context, channelID *int64, status *bool, userID int64) (int64, error)
	CreateChannelConfig(ctx context.Context, countryCodes []string, ids []int64, userID int64) (int64, error)
	UpdateChannelConfig(ctx context.Context, countryCodes []string, ids []int64, status bool, userID int64) (int, error)

	GetChannelRecommends(ctx context.Context,
		countryCode, sorter *string, channelID *int64,
		status, isRecommended *bool,
		current, pageSize int) ([]*domainEntity.ChannelRecommend, error)
	CountChannelRecommends(ctx context.Context, countryCode *string,
		channelID *int64, status, isRecommended *bool) (int64, error)
	CreateChannelRecommend(ctx context.Context, countryCode string, channelID int64) (int64, error)
	UpdateChannelRecommend(ctx context.Context, countryCode string, channelID int64,
		isRecommended, status bool) (int, error)
}
