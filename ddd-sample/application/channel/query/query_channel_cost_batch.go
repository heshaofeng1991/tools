/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_cost_batch
	@Date    2022/5/16 17:24
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
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelInfoHandler(quoteRepo domainRepository.ChannelRepository) GetChannelInfoHandler {
	return GetChannelInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelInfoHandler) Handle(
	ctx context.Context,
	channelPlatform, channelCode, channelName, sorter *string,
	current, pageSize int,
) ([]*domainEntity.Channel, int64, error) {
	result, err := q.quoteRepo.GetChannels(
		ctx,
		channelPlatform,
		channelCode,
		channelName,
		sorter,
		current,
		pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	total, err := q.quoteRepo.CountChannels(
		ctx,
		channelPlatform,
		channelCode,
		channelName)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	return result, total, nil
}
