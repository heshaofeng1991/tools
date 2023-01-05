/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_cost_batch_by_id
	@Date    2022/5/17 14:33
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelCostBatchByIDHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelCostBatchByIDHandler(quoteRepo domainRepository.ChannelRepository) GetChannelCostBatchByIDHandler {
	return GetChannelCostBatchByIDHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostBatchByIDHandler) Handle(
	ctx context.Context,
	id int64,
) (*domainEntity.ChannelCostBatch, error) {
	result, err := q.quoteRepo.GetChannelCostBatchByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
