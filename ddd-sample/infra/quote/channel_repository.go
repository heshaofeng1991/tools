/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/12 10:08
	@Desc
*/

package quote

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/pkg/errors"
)

type ShippingOptionRepository struct {
	entClient *ent.Client
}

func NewShippingOptionRepository(entClient *ent.Client) *ShippingOptionRepository {
	return &ShippingOptionRepository{entClient: entClient}
}

func (c ShippingOptionRepository) GetChannels(ctx context.Context, prepayTariff bool,
	presetChannelIds, testChannelIds, excludeChannelIds []int64, warehouseID *int64,
) ([]*domainEntity.Channel, error) {
	result, err := c.getAvailableChannels(ctx, prepayTariff,
		presetChannelIds, testChannelIds, excludeChannelIds, warehouseID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return domainEntity.CovertDBToChannel(result), nil
}

// 获取可用的渠道.
func (c ShippingOptionRepository) getAvailableChannels(ctx context.Context, prepayTariff bool,
	presetChannelIds, testChannelIds, excludeChannelIds []int64, warehouseID *int64,
) (rsp []*ent.Channel, err error) {
	// 	组装查询渠道信息的查询条件.
	query := c.getAvailableChannelsCondition(prepayTariff, presetChannelIds,
		testChannelIds, excludeChannelIds, warehouseID)

	rsp, err = query.All(ctx)

	return rsp, errors.Wrap(err, "")
}

// 组装获取可用渠道的查询条件.
func (c ShippingOptionRepository) getAvailableChannelsCondition(prepayTariff bool, presetChannelIds,
	testChannelIds, excludeChannelIds []int64, warehouseID *int64,
) (query *ent.ChannelQuery) {
	var (
		lenPresetChannelIds, lenTestChannelIds int
		channelIds                             []int64
	)

	lenPresetChannelIds = len(presetChannelIds)
	lenTestChannelIds = len(testChannelIds)

	query = c.entClient.Channel.Query().Where(channel.Status(1))

	if len(excludeChannelIds) > 0 {
		query.Where(channel.IDNotIn(excludeChannelIds...))
	}

	if warehouseID != nil {
		query.Where(channel.WarehouseIDEQ(*warehouseID))
	}

	switch {
	case lenPresetChannelIds == 0 && lenTestChannelIds > 0:
		query = query.Where(
			channel.Or(
				channel.Test(false),
				channel.And(
					channel.Test(true),
					channel.IDIn(testChannelIds...),
				),
			),
		)
	case lenPresetChannelIds > 0:
		channelIds = make([]int64, 0)
		channelIds = append(channelIds, presetChannelIds...)

		if lenTestChannelIds > 0 {
			channelIds = append(channelIds, testChannelIds...)
		}

		query = query.Where(channel.IDIn(channelIds...))
	default:
		query = query.Where(channel.Test(false))
	}

	if prepayTariff {
		query = query.Where(channel.PrepayTariff(true))
	}

	return query
}
