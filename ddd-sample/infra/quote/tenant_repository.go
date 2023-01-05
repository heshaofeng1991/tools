/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    tenant
	@Date    2022/5/12 10:09
	@Desc
*/

package quote

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	ent "github.com/heshaofeng1991/entgo/ent/gen"
	"github.com/heshaofeng1991/entgo/ent/gen/user"
	"github.com/heshaofeng1991/entgo/ent/viewer"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetTenantInfo(ctx context.Context, userID int64) (*domainEntity.Tenant, error) {
	ctx = viewer.NewContext(ctx, viewer.UserViewer{T: &ent.Tenant{ID: -1}})

	tenantInfo, err := c.entClient.User.Query().Where(user.IDEQ(userID)).QueryTenant().First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	prepayTariff := false

	if tenantInfo.PrepayTariff == 1 {
		prepayTariff = true
	}

	return domainEntity.NewTenant(
		tenantInfo.ID,
		prepayTariff,
		tenantInfo.PresetChannelIds,
		tenantInfo.TestChannelIds), nil
}
