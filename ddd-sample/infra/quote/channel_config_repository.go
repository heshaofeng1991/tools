/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_config_repository
	@Date    2022/5/25 18:33
	@Desc
*/

package quote

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/customerconfig"
	"github.com/heshaofeng1991/entgo/ent/gen/user"
	"github.com/heshaofeng1991/entgo/ent/viewer"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetChannelConfigsByIDs(ctx context.Context, ids []int64,
	countryCode string, userID int64,
) ([]*domainEntity.ChannelConfig, error) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{T: &ent.Tenant{ID: -1}})

	tenant, err := c.entClient.User.Query().Where(user.IDEQ(userID)).QueryTenant().First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	ctx = viewer.NewContext(ctx, viewer.UserViewer{T: &ent.Tenant{ID: tenant.ID}})

	qry := c.entClient.CustomerConfig.Query().Where(
		customerconfig.And(
			customerconfig.ChannelIDIn(ids...),
			customerconfig.StatusEQ(1),
			customerconfig.DeletedAtIsNil(),
		))
	qry = qry.Where(func(s *sql.Selector) {
		s.Where(sqljson.ValueContains(customerconfig.FieldExcludeCountryCode, countryCode))
	})

	result, err := qry.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return domainEntity.CovertDBToChannelConfig(result), errors.Wrap(err, "")
}
