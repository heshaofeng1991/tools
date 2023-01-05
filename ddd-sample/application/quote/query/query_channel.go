/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel
	@Date    2022/5/11 18:00
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelInfoHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetChannelInfoHandler(quoteRepo domainRepository.Repository) GetChannelInfoHandler {
	return GetChannelInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelInfoHandler) Handle(ctx context.Context, prepayTariff bool,
	presetChannelIds, testChannelIds, excludeChannelIds []int64, warehouseID *int64,
) ([]*domainEntity.Channel, error) {
	result, err := q.quoteRepo.GetChannels(ctx, prepayTariff, presetChannelIds,
		testChannelIds, excludeChannelIds, warehouseID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
