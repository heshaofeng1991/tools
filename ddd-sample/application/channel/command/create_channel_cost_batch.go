/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel_cost_batch
	@Date    2022/5/16 19:49
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelCostBatchHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelCostBatchHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelCostBatchHandler {
	return CreateChannelCostBatchHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelCostBatchHandler) Handle(ctx context.Context, effectiveDate string,
	channelID int64, status bool,
) (int64, error) {
	result, err := q.quoteRepo.CreateChannelCostBatch(
		ctx,
		effectiveDate,
		channelID,
		status)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
