/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_cost
	@Date    2022/5/16 22:02
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelCostInfoHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelCostInfoHandler(quoteRepo domainRepository.ChannelRepository) GetChannelCostInfoHandler {
	return GetChannelCostInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostInfoHandler) Handle(
	ctx context.Context, countryCode *string,
	channelID, channelCostBatchID int64,
	current, pageSize int,
) ([]*domainEntity.ChannelCost, int64, error) {
	result, err := q.quoteRepo.GetChannelCosts(
		ctx,
		channelID,
		channelCostBatchID,
		countryCode,
		current,
		pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	total, err := q.quoteRepo.CountChannelCosts(
		ctx,
		channelID,
		channelCostBatchID,
		countryCode)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	return result, total, nil
}
