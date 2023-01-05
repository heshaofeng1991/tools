/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend_repository
	@Date    2022/5/25 15:17
	@Desc
*/

package channel

import (
	"context"
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channelrecommend"
	"github.com/pkg/errors"
)

func (c ChlRepository) GetChannelRecommends(ctx context.Context, countryCode, sorter *string,
	channelID *int64, status, isRecommended *bool,
	current, pageSize int,
) ([]*domainEntity.ChannelRecommend, error) {
	qry := c.entClient.ChannelRecommend.Query().Where(channelrecommend.DeletedAtIsNil())

	if countryCode != nil {
		qry.Where(channelrecommend.CountryCodeEQ(*countryCode))
	}

	if channelID != nil {
		qry.Where(channelrecommend.ChannelIDEQ(*channelID))
	}

	if status != nil && *status {
		qry.Where(channelrecommend.StatusEQ(1))
	}

	if status != nil && !*status {
		qry.Where(channelrecommend.StatusEQ(0))
	}

	if isRecommended != nil && *isRecommended {
		qry.Where(channelrecommend.IsRecommended(1))
	}

	if isRecommended != nil && !*isRecommended {
		qry.Where(channelrecommend.IsRecommended(0))
	}

	var newSort string

	if sorter != nil {
		newSort = strings.ReplaceAll(*sorter, "\\", "")
		if !strings.Contains(newSort, "desc") && !strings.Contains(newSort, "asc") {
			return nil, errors.New("invalid sort order")
		}

		if strings.Contains(newSort, "asc") {
			qry.Order(ent.Asc(channelrecommend.FieldUpdatedAt))
		}

		qry.Order(ent.Desc(channelrecommend.FieldUpdatedAt))
	}

	if sorter == nil {
		qry.Order(ent.Desc(channelrecommend.FieldUpdatedAt))
	}

	results, err := qry.Limit(pageSize).Offset((current - 1) * pageSize).All(ctx)

	return domainEntity.CovertDBToChannelRecommend(results), errors.Wrap(err, "")
}

func (c ChlRepository) CountChannelRecommends(ctx context.Context, countryCode *string,
	channelID *int64, status, isRecommended *bool,
) (int64, error) {
	qry := c.entClient.ChannelRecommend.Query().Where(channelrecommend.DeletedAtIsNil())

	if countryCode != nil {
		qry.Where(channelrecommend.CountryCodeEQ(*countryCode))
	}

	if channelID != nil {
		qry.Where(channelrecommend.ChannelIDEQ(*channelID))
	}

	if status != nil && *status {
		qry.Where(channelrecommend.StatusEQ(1))
	}

	if status != nil && !*status {
		qry.Where(channelrecommend.StatusEQ(0))
	}

	if isRecommended != nil && *isRecommended {
		qry.Where(channelrecommend.IsRecommended(1))
	}

	if isRecommended != nil && !*isRecommended {
		qry.Where(channelrecommend.IsRecommended(0))
	}

	result, err := qry.Count(ctx)

	return int64(result), err
}

func (c ChlRepository) CreateChannelRecommend(ctx context.Context, countryCode string, channelID int64) (int64, error) {
	existed, err := c.entClient.ChannelRecommend.Query().Where(
		channelrecommend.And(
			channelrecommend.DeletedAtIsNil(),
			channelrecommend.CountryCodeEQ(countryCode))).
		Exist(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	if existed {
		return 0, errors.New("channel has already been created")
	}

	if _, err := c.entClient.ChannelRecommend.Create().
		SetChannelID(channelID).
		SetCountryCode(countryCode).
		Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 0, nil
}

func (c ChlRepository) UpdateChannelRecommend(ctx context.Context, countryCode string,
	channelID int64, isRecommended, status bool,
) (int, error) {
	var sts, isRd int8

	if isRecommended {
		isRd = 1
	}

	if status {
		sts = 1
	}

	if _, err := c.entClient.ChannelRecommend.Update().
		SetIsRecommended(isRd).
		SetStatus(sts).
		Where(
			channelrecommend.And(
				channelrecommend.ChannelID(channelID),
				channelrecommend.CountryCode(countryCode))).
		Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 0, nil
}
