/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_config_repository
	@Date    2022/5/25 15:16
	@Desc
*/

package channel

import (
	"context"
	"encoding/json"
	"strings"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/customerconfig"
	"github.com/heshaofeng1991/entgo/ent/gen/user"
	"github.com/heshaofeng1991/entgo/ent/viewer"
	"github.com/pkg/errors"
)

func (c ChlRepository) getTenantInfo(ctx context.Context, userID int64) (context.Context, *ent.Tenant, error) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{T: &ent.Tenant{ID: -1}})

	tenant, err := c.entClient.User.Query().Where(user.IDEQ(userID)).QueryTenant().First(ctx)
	if err != nil {
		return ctx, nil, errors.Wrap(err, "")
	}

	ctx = viewer.NewContext(ctx, viewer.UserViewer{T: &ent.Tenant{ID: tenant.ID}})

	return ctx, tenant, nil
}

func (c ChlRepository) GetChannelConfigs(ctx context.Context, sorter *string, channelID *int64,
	status *bool, current, pageSize int, userID int64,
) ([]*domainEntity.ChannelConfig, error) {
	ctx, _, err := c.getTenantInfo(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	qry := c.entClient.CustomerConfig.Query().Where(customerconfig.DeletedAtIsNil())

	if channelID != nil {
		qry.Where(customerconfig.ChannelIDEQ(*channelID))
	}

	if status != nil && *status {
		qry.Where(customerconfig.StatusEQ(1))
	}

	if status != nil && !*status {
		qry.Where(customerconfig.StatusEQ(0))
	}

	var newSort string

	if sorter != nil {
		newSort = strings.ReplaceAll(*sorter, "\\", "")
		if !strings.Contains(newSort, "desc") && !strings.Contains(newSort, "asc") {
			return nil, errors.New("invalid sort order")
		}

		if strings.Contains(newSort, "asc") {
			qry.Order(ent.Asc(customerconfig.FieldUpdatedAt))
		}

		qry.Order(ent.Desc(customerconfig.FieldUpdatedAt))
	}

	if sorter == nil {
		qry.Order(ent.Desc(customerconfig.FieldUpdatedAt))
	}

	results, err := qry.Limit(pageSize).Offset((current - 1) * pageSize).All(ctx)

	return domainEntity.CovertDBToChannelConfig(results), errors.Wrap(err, "")
}

func (c ChlRepository) CountChannelConfigs(ctx context.Context,
	channelID *int64, status *bool, userID int64,
) (int64, error) {
	ctx, _, err := c.getTenantInfo(ctx, userID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	qry := c.entClient.CustomerConfig.Query().Where(customerconfig.DeletedAtIsNil())

	if channelID != nil {
		qry.Where(customerconfig.ChannelIDEQ(*channelID))
	}

	if status != nil && *status {
		qry.Where(customerconfig.StatusEQ(1))
	}

	if status != nil && !*status {
		qry.Where(customerconfig.StatusEQ(0))
	}

	total, err := qry.Count(ctx)

	return int64(total), errors.Wrap(err, "")
}

func (c ChlRepository) CreateChannelConfig(ctx context.Context,
	countryCodes []string, ids []int64, userID int64,
) (int64, error) {
	ctx, tenant, err := c.getTenantInfo(ctx, userID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	existed, err := c.entClient.CustomerConfig.Query().Where(
		customerconfig.And(
			customerconfig.ChannelIDIn(ids...),
			customerconfig.DeletedAtIsNil())).Exist(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	if existed {
		return 0, errors.New("user has already been created channel config data")
	}

	countryCode, _ := json.Marshal(countryCodes) //nolint:errchkjson

	customerConfigCreate := make([]*ent.CustomerConfigCreate, 0)

	for _, id := range ids {
		val := c.entClient.CustomerConfig.Create().
			SetTenant(tenant).
			SetChannelID(id).
			SetExcludeCountryCode(string(countryCode)).
			SetStatus(1)

		customerConfigCreate = append(customerConfigCreate, val)
	}

	if _, err := c.entClient.CustomerConfig.CreateBulk(customerConfigCreate...).Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}

func (c ChlRepository) UpdateChannelConfig(ctx context.Context, countryCodes []string,
	ids []int64, status bool, userID int64,
) (int, error) {
	ctx, tenant, err := c.getTenantInfo(ctx, userID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	countryCode, _ := json.Marshal(countryCodes) //nolint:errchkjson

	var flag int8
	if status {
		flag = 1
	}

	if _, err := c.entClient.CustomerConfig.Update().
		SetTenant(tenant).
		SetExcludeCountryCode(string(countryCode)).
		SetStatus(flag).
		Where(customerconfig.ChannelIDIn(ids...)).
		Save(ctx); err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}
