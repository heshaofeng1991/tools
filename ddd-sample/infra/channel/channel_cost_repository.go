/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
	@Date    2022/5/12 10:08
	@Desc
*/

package channel

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcost"
	"github.com/pkg/errors"
)

func (c ChlRepository) GetChannelCosts(ctx context.Context,
	channelID, channelCostBatchID int64,
	countryCode *string, current, pageSize int,
) ([]*domainEntity.ChannelCost, error) {
	qry := c.entClient.ChannelCost.Query().Where(
		channelcost.And(
			channelcost.DeletedAtIsNil(),
			channelcost.ChannelIDEQ(channelID),
			channelcost.ChannelCostBatchIDEQ(channelCostBatchID)))

	if countryCode != nil {
		qry.Where(channelcost.CountryCodeEQ(*countryCode))
	}

	results, err := qry.WithChannels(func(query *ent.ChannelQuery) {
		query.Where(channel.DeletedAtIsNil())
	}).Order(ent.Desc(channelcost.FieldUpdatedAt)).
		Limit(pageSize).Offset((current - 1) * pageSize).All(ctx)

	return domainEntity.CovertChannelCost(results), errors.Wrap(err, "")
}

func (c ChlRepository) CountChannelCosts(ctx context.Context,
	channelID, channelCostBatchID int64, countryCode *string,
) (int64, error) {
	qry := c.entClient.ChannelCost.Query().Where(
		channelcost.And(
			channelcost.DeletedAtIsNil(),
			channelcost.ChannelIDEQ(channelID),
			channelcost.ChannelCostBatchIDEQ(channelCostBatchID)))

	if countryCode != nil {
		qry.Where(channelcost.CountryCodeEQ(*countryCode))
	}

	total, err := qry.Count(ctx)

	return int64(total), errors.Wrap(err, "")
}

func (c ChlRepository) CreateChannelCost(ctx context.Context, costs []*domainEntity.ChannelCost) error {
	channelCosts := make([]*ent.ChannelCostCreate, 0)

	for _, val := range costs {
		channelCosts = append(channelCosts, c.BuildChannelCost(val))
	}

	_, err := c.entClient.ChannelCost.CreateBulk(channelCosts...).Save(ctx)

	return errors.Wrap(err, "")
}

func (c ChlRepository) BuildChannelCost(cost *domainEntity.ChannelCost) *ent.ChannelCostCreate {
	entCreate := c.entClient.ChannelCost.Create()

	if cost.Zone() != "" {
		entCreate.SetZone(cost.Zone())
	}

	return entCreate.SetChannelCostBatchID(cost.ChannelCostBatchID()).
		SetChannelID(cost.ChannelID()).
		SetMode(cost.Mode()).
		SetCountryCode(cost.CountryCode()).
		SetStartWeight(cost.StartWeight()).
		SetEndWeight(cost.EndWeight()).
		SetFirstWeight(cost.FirstWeight()).
		SetFirstWeightFee(cost.FirstWeightFee()).
		SetUnitWeight(cost.UnitWeight()).
		SetUnitWeightFee(cost.UnitWeightFee()).
		SetFuelFee(cost.FuelFee()).
		SetProcessingFee(cost.ProcessingFee()).
		SetRegistrationFee(cost.RegistrationFee()).
		SetMiscFee(cost.MiscFee()).
		SetMinNormalDays(cost.MinNormalDays()).
		SetMaxNormalDays(cost.MaxNormalDays()).
		SetAverageDays(cost.AverageDays())
}

func (c ChlRepository) UpdateChannelCostStatus(ctx context.Context, ids []int64,
	countryCodes []string, status bool,
) (int, error) {
	var sts int8

	if status {
		sts = 1
	}

	if _, err := c.entClient.ChannelCost.Update().
		SetStatus(sts).
		Where(channelcost.And(
			channelcost.ChannelIDIn(ids...),
			channelcost.CountryCodeIn(countryCodes...),
			channelcost.DeletedAtIsNil())).Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}
