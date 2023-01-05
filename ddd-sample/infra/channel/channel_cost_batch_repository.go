/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
	@Date    2022/5/12 10:08
	@Desc
*/

package channel

import (
	"context"
	"strings"
	"time"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/channelcostbatche"
	"github.com/pkg/errors"
)

const (
	UTC = "2006-01-02T15:04:05Z"
)

func (c ChlRepository) GetChannelCostBatches(ctx context.Context, effectiveDate, sorter *string,
	channelID int64, status *bool, current, pageSize int,
) ([]*domainEntity.ChannelCostBatch, error) {
	qry := c.entClient.ChannelCostBatche.Query().Where(
		channelcostbatche.And(
			channelcostbatche.ChannelIDEQ(channelID),
			// channelcostbatche.DeletedAtIsNil(),
		),
	)

	if effectiveDate != nil {
		timeRes, err := time.Parse(UTC, *effectiveDate)
		if err != nil {
			return nil, errors.Wrap(err, "")
		}

		qry.Where(channelcostbatche.EffectiveDateEQ(timeRes))
	}

	if status != nil {
		qry.Where(channelcostbatche.StatusEQ(*status))
	}

	var newSort string

	if sorter != nil {
		newSort = strings.ReplaceAll(*sorter, "\\", "")
		if !strings.Contains(newSort, "desc") && !strings.Contains(newSort, "asc") {
			return nil, errors.New("invalid sort order")
		}

		if strings.Contains(newSort, "asc") {
			qry.Order(ent.Asc(channelcostbatche.FieldUpdatedAt))
		}

		qry.Order(ent.Desc(channelcostbatche.FieldUpdatedAt))
	}

	if sorter == nil {
		qry.Order(ent.Desc(channelcostbatche.FieldUpdatedAt))
	}

	results, err := qry.Limit(pageSize).Offset((current - 1) * pageSize).All(ctx)

	return domainEntity.CovertDBToChannelCostBatch(results), errors.Wrap(err, "")
}

func (c ChlRepository) CountChannelCostBatches(ctx context.Context, effectiveDate *string,
	channelID int64, status *bool,
) (int64, error) {
	qry := c.entClient.ChannelCostBatche.Query().Where(
		channelcostbatche.And(
			channelcostbatche.ChannelIDEQ(channelID),
			// channelcostbatche.DeletedAtIsNil(),
		),
	)

	if effectiveDate != nil {
		timeRes, err := time.Parse(UTC, *effectiveDate)
		if err != nil {
			return 0, errors.Wrap(err, "")
		}

		qry.Where(channelcostbatche.EffectiveDateEQ(timeRes))
	}

	if status != nil {
		qry.Where(channelcostbatche.StatusEQ(*status))
	}

	total, err := qry.Count(ctx)

	return int64(total), errors.Wrap(err, "")
}

func (c ChlRepository) CreateChannelCostBatch(ctx context.Context, effectiveDate string,
	channelID int64, status bool,
) (int64, error) {
	timeRes, err := time.Parse(UTC, effectiveDate)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	result, err := c.entClient.ChannelCostBatche.
		Create().
		SetChannelID(channelID).
		SetStatus(status).
		SetEffectiveDate(timeRes).
		Save(ctx)

	return result.ID, errors.Wrap(err, "")
}

func (c ChlRepository) UpdateChannelCostBatch(ctx context.Context, effectiveDate string,
	channelID, id int64, status bool,
) (int, error) {
	timeRes, err := time.Parse(UTC, effectiveDate)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	if _, err := c.entClient.ChannelCostBatche.
		UpdateOneID(id).
		SetChannelID(channelID).
		SetStatus(status).
		SetEffectiveDate(timeRes).
		Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, errors.Wrap(err, "")
}

func (c ChlRepository) GetChannelCostBatchByID(ctx context.Context, id int64) (*domainEntity.ChannelCostBatch, error) {
	result, err := c.entClient.ChannelCostBatche.Get(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return domainEntity.NewChannelCostBatch(
		result.ID,
		result.ChannelID,
		result.EffectiveDate,
		result.ExpiryDate,
		result.Status,
		result.CreatedAt,
		result.UpdatedAt), nil
}
