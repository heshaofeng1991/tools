/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_config
	@Date    2022/5/25 22:50
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelConfigHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetChannelConfigHandler(quoteRepo domainRepository.Repository) GetChannelConfigHandler {
	return GetChannelConfigHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelConfigHandler) Handle(
	ctx context.Context, ids []int64,
	countryCode string, userID int64,
) ([]*domainEntity.ChannelConfig, error) {
	result, err := q.quoteRepo.GetChannelConfigsByIDs(
		ctx,
		ids,
		countryCode,
		userID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
