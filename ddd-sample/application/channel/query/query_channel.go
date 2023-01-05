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

type GetChannelCostBatchInfoHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelCostBatchInfoHandler(quoteRepo domainRepository.ChannelRepository) GetChannelCostBatchInfoHandler {
	return GetChannelCostBatchInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostBatchInfoHandler) Handle(
	ctx context.Context, effectiveDate, sorter *string,
	channelID int64, status *bool,
	current, pageSize int,
) ([]*domainEntity.ChannelCostBatch, int64, error) {
	result, err := q.quoteRepo.GetChannelCostBatches(
		ctx,
		effectiveDate,
		sorter,
		channelID,
		status,
		current,
		pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	total, err := q.quoteRepo.CountChannelCostBatches(
		ctx,
		effectiveDate,
		channelID,
		status)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	return result, total, nil
}
