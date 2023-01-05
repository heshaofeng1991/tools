/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    shipping_option_repository
	@Date    2022/5/15 17:55
	@Desc
*/

package repository

import (
	"context"
	"time"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
)

type Repository interface {
	GetChannels(ctx context.Context, prepayTariff bool, presetChannelIds,
		testChannelIds, excludeChannelIds []int64, warehouseID *int64) ([]*domainEntity.Channel, error)
	GetCountryZone(ctx context.Context, zipCode, countryCode string) ([]*domainEntity.CountryZone, error)
	GetChannelCostBatch(ctx context.Context, ids []int64, date time.Time) ([]*domainEntity.ChannelCostBatch, error)
	GetChannelCostByWeight(ctx context.Context, batchIDs []int64,
		countryCode string, actualGrams int) ([]*domainEntity.ChannelCost, error)
	GetChannelCostByVolume(ctx context.Context, countryCode string, actualGrams int, volume int64,
		volumeChannelCostBatches []*domainEntity.ChannelBatch) (rsp []*domainEntity.ChannelCostRsp, err error)
	GetVatInfoByCountryCode(ctx context.Context, countryCode string) (*domainEntity.Vat, error)
	GetTenantInfo(ctx context.Context, userID int64) (*domainEntity.Tenant, error)
	GetChannelRecommendsByCondition(ctx context.Context, ids []int64,
		countryCode string) (*domainEntity.ChannelRecommend, error)
	GetChannelConfigsByIDs(ctx context.Context, ids []int64,
		countryCode string, userID int64) ([]*domainEntity.ChannelConfig, error)
}
