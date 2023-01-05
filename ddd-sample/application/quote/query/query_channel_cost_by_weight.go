/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost
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

type GetChannelCostInfoByWeightHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetChannelCostInfoByWeightHandler(quoteRepo domainRepository.Repository) GetChannelCostInfoByWeightHandler {
	return GetChannelCostInfoByWeightHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostInfoByWeightHandler) Handle(ctx context.Context, batchIDs []int64,
	countryCode string, actualGrams int,
) ([]*domainEntity.ChannelCost, error) {
	result, err := q.quoteRepo.GetChannelCostByWeight(ctx, batchIDs, countryCode, actualGrams)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
