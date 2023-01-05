/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel_cost_batch
	@Date    2022/5/16 19:44
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelCostBatchHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateChannelCostBatchHandler(quoteRepo domainRepository.ChannelRepository) UpdateChannelCostBatchHandler {
	return UpdateChannelCostBatchHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelCostBatchHandler) Handle(ctx context.Context, effectiveDate string,
	channelID, id int64, status bool,
) (int, error) {
	result, err := q.quoteRepo.UpdateChannelCostBatch(
		ctx,
		effectiveDate,
		channelID,
		id,
		status)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return result, nil
}
