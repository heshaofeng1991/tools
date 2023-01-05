/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/12 10:08
	@Desc
*/

package channel

import (
	"context"
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/pkg/errors"
)

type ChlRepository struct {
	entClient *ent.Client
}

func NewChlRepository(entClient *ent.Client) *ChlRepository {
	return &ChlRepository{
		entClient: entClient,
	}
}

func (c ChlRepository) GetChannels(
	ctx context.Context,
	channelPlatform, channelCode, channelName, sorter *string,
	current, pageSize int,
) ([]*domainEntity.Channel, error) {
	qry := c.entClient.Channel.Query().Where(channel.DeletedAtIsNil())

	if channelPlatform != nil {
		qry.Where(channel.CourierPlatformEQ(*channelPlatform))
	}

	if channelCode != nil {
		qry.Where(channel.CodeEQ(*channelCode))
	}

	if channelName != nil {
		qry.Where(channel.NameEQ(*channelName))
	}

	var newSort string

	if sorter != nil {
		newSort = strings.ReplaceAll(*sorter, "\\", "")
		if !strings.Contains(newSort, "desc") && !strings.Contains(newSort, "asc") {
			return nil, errors.New("invalid sort order")
		}

		if strings.Contains(newSort, "asc") {
			qry.Order(ent.Asc(channel.FieldUpdatedAt))
		}

		qry.Order(ent.Desc(channel.FieldUpdatedAt))
	}

	if sorter == nil {
		qry.Order(ent.Desc(channel.FieldUpdatedAt))
	}

	results, err := qry.Limit(pageSize).Offset((current - 1) * pageSize).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return domainEntity.CovertDBToChannel(results), nil
}

func (c ChlRepository) CountChannels(
	ctx context.Context,
	channelPlatform, channelCode, channelName *string,
) (int64, error) {
	qry := c.entClient.Channel.Query().Where(channel.DeletedAtIsNil())

	if channelPlatform != nil {
		qry.Where(channel.CourierPlatformEQ(*channelPlatform))
	}

	if channelCode != nil {
		qry.Where(channel.CodeEQ(*channelCode))
	}

	if channelName != nil {
		qry.Where(channel.NameEQ(*channelName))
	}

	total, err := qry.Count(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return int64(total), nil
}

func (c ChlRepository) CreateChannel(ctx context.Context, channel domainEntity.Channel) (int64, error) {
	result, err := c.entClient.Channel.Create().
		SetWarehouseID(channel.WarehouseID()).
		SetCourierPlatform(channel.CourierPlatform()).
		SetName(channel.Name()).
		SetCode(channel.Code()).
		SetType(channel.ChannelLogisticType()).
		SetQuotationCurrency(channel.QuotationCurrency()).
		SetVolumeFactor(channel.VolumeFactor()).
		SetEnName(channel.EnName()).
		SetDisplayName(channel.DisplayName()).
		SetHasTrackingNumber(channel.HasTrackingNumber()).
		SetMaxLength(channel.MaxLength()).
		SetMinLength(channel.MinLength()).
		SetMaxThreeSideSum(channel.MaxThreeSideSum()).
		SetDescription(channel.Description()).
		SetSortingPort(channel.SortingPort()).
		SetPrepayTariff(channel.PrepayTariff()).
		SetStatus(channel.Status()).
		SetTest(channel.Test()).
		SetExcludeAttributes(channel.ExcludeAttributes()).
		SetOptions(channel.Options()).
		SetVirtual(channel.Virtual()).
		SetChannelType(channel.ChannelType()).
		SetDeliverDuty(channel.DeliverDuty()).
		SetSpecial(channel.Special()).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result.ID, nil
}

func (c ChlRepository) UpdateChannel(ctx context.Context, channel domainEntity.Channel) (int, error) {
	_, err := c.entClient.Channel.UpdateOneID(channel.ID()).
		SetWarehouseID(channel.WarehouseID()).
		SetCourierPlatform(channel.CourierPlatform()).
		SetName(channel.Name()).
		SetCode(channel.Code()).
		SetType(channel.ChannelLogisticType()).
		SetQuotationCurrency(channel.QuotationCurrency()).
		SetVolumeFactor(channel.VolumeFactor()).
		SetEnName(channel.EnName()).
		SetDisplayName(channel.DisplayName()).
		SetHasTrackingNumber(channel.HasTrackingNumber()).
		SetMaxLength(channel.MaxLength()).
		SetMinLength(channel.MinLength()).
		SetMaxThreeSideSum(channel.MaxThreeSideSum()).
		SetDescription(channel.Description()).
		SetSortingPort(channel.SortingPort()).
		SetPrepayTariff(channel.PrepayTariff()).
		SetStatus(channel.Status()).
		SetTest(channel.Test()).
		SetExcludeAttributes(channel.ExcludeAttributes()).
		SetOptions(channel.Options()).
		SetChannelType(channel.ChannelType()).
		SetVirtual(channel.Virtual()).
		SetDeliverDuty(channel.DeliverDuty()).
		SetSpecial(channel.Special()).
		Save(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}
