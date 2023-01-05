/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    channel_cost_batch
	@Date    2022/5/11 18:01
	@Desc
*/

package query

import (
	"context"
	"time"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelCostBatchInfoHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetChannelCostBatchInfoHandler(quoteRepo domainRepository.Repository) GetChannelCostBatchInfoHandler {
	return GetChannelCostBatchInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostBatchInfoHandler) Handle(ctx context.Context,
	ids []int64, date time.Time,
) ([]*domainEntity.ChannelCostBatch, error) {
	result, err := q.quoteRepo.GetChannelCostBatch(ctx, ids, date)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
