/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_attribute_repository
	@Date    2022/5/20 12:19
	@Desc
*/

package channel

import (
	"context"
	"encoding/json"

	"github.com/heshaofeng1991/common/dao"
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/attribute"
	"github.com/heshaofeng1991/entgo/ent/gen/channel"
	"github.com/pkg/errors"
)

func (c ChlRepository) GetChannelAttributes(ctx context.Context) ([]*domainEntity.ChannelAttribute, error) {
	results, err := c.entClient.Attribute.Query().Where(
		attribute.And(
			attribute.TypeEQ(3)),
		attribute.Status(1),
		attribute.DeletedAtIsNil(),
	).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	rsp := make([]*domainEntity.ChannelAttribute, 0)

	for _, val := range results {
		rsp = append(rsp, domainEntity.NewChannelAttribute(val.Value))
	}

	return rsp, nil
}

func (c ChlRepository) CreateChannelAttributes(ctx context.Context, attribute string,
	updateFn func(ctx context.Context, attribute string) (int32, error),
) (int32, error) {
	cls, err := c.entClient.Channel.Query().Where(
		channel.And(
			channel.IDGT(0),
			channel.DeletedAtIsNil())).All(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "query channel attributes failed")
	}

	if err := dao.WithTx(
		ctx,
		c.entClient,
		func(transaction *ent.Tx) error {
			if updateFn != nil {
				if _, err := updateFn(ctx, attribute); err != nil {
					return err
				}
			}

			if _, err := transaction.Attribute.Create().
				SetType(3).
				SetValue(attribute).
				Save(ctx); err != nil {
				return errors.Wrap(err, "create channel attribute failed")
			}

			for _, val := range cls {
				if val.ExcludeAttributes == "" || val.ExcludeAttributes == "[]" {
					continue
				}

				attributes := make([]string, 0)
				newAttributes := make([]string, 0)
				_ = json.Unmarshal([]byte(val.ExcludeAttributes), &attributes)
				for _, attr := range attributes {
					if attr != attribute {
						newAttributes = append(newAttributes, attr)
					}
				}

				setAttributes, _ := json.Marshal(newAttributes)
				_, err = transaction.Channel.Update().
					SetExcludeAttributes(string(setAttributes)).
					Where(channel.IDEQ(val.ID)).
					Save(ctx)
				if err != nil {
					break
				}
			}

			return errors.Wrap(err, "update channel attributes failed")
		},
	); err != nil {
		return 0, errors.Wrap(err, "create channel attribute failed")
	}

	return 1, nil
}

func (c ChlRepository) UpdateChannelAttributes(ctx context.Context, attribute string,
	channelIDs []int64,
) ([]*domainEntity.UpdateChannelAttribute, error) {
	cls, err := c.entClient.Channel.Query().Where(
		channel.And(
			channel.IDIn(channelIDs...),
			channel.DeletedAtIsNil())).All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	rsp := make([]*domainEntity.UpdateChannelAttribute, 0)

	for _, val := range cls {
		if val.ExcludeAttributes == "" || val.ExcludeAttributes == "[]" {
			continue
		}

		attributes := make([]string, 0)
		newAttributes := make([]string, 0)
		_ = json.Unmarshal([]byte(val.ExcludeAttributes), &attributes)

		for _, attr := range attributes {
			if attr != attribute {
				newAttributes = append(newAttributes, attr)
			}
		}

		setAttributes, _ := json.Marshal(newAttributes) //nolint:errchkjson
		if _, err = c.entClient.Channel.Update().
			SetExcludeAttributes(string(setAttributes)).
			Where(channel.IDEQ(val.ID)).
			Save(ctx); err != nil {
			rsp = append(rsp, domainEntity.NewUpdateChannelAttribute(val.ID, 0, err.Error()))

			continue
		}

		rsp = append(rsp, domainEntity.NewUpdateChannelAttribute(val.ID, 1, ""))
	}

	return rsp, nil
}
