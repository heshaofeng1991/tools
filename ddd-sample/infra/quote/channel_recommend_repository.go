/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_recommend_repository
	@Date    2022/5/25 18:32
	@Desc
*/

package quote

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	"github.com/heshaofeng1991/entgo/ent/gen/channelrecommend"
	"github.com/pkg/errors"
)

func (c ShippingOptionRepository) GetChannelRecommendsByCondition(ctx context.Context, ids []int64,
	countryCode string,
) (*domainEntity.ChannelRecommend, error) {
	result, err := c.entClient.ChannelRecommend.Query().Where(
		channelrecommend.And(
			channelrecommend.ChannelIDIn(ids...),
			channelrecommend.StatusEQ(1),
			channelrecommend.CountryCode(countryCode),
			channelrecommend.DeletedAtIsNil(),
		)).First(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return domainEntity.NewChannelRecommend(
		result.ID,
		result.CreatedAt,
		result.UpdatedAt,
		result.DeletedAt,
		result.CountryCode,
		result.ChannelID,
		result.IsRecommended,
		result.Status,
		result.Value), errors.Wrap(err, "")
}
