/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_config
	@Date    2022/5/25 12:39
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
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelConfigHandler(quoteRepo domainRepository.ChannelRepository) GetChannelConfigHandler {
	return GetChannelConfigHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelConfigHandler) Handle(
	ctx context.Context,
	sorter *string,
	channelID *int64,
	status *bool,
	current, pageSize int,
	userID int64,
) ([]*domainEntity.ChannelConfig, int64, error) {
	result, err := q.quoteRepo.GetChannelConfigs(
		ctx,
		sorter,
		channelID,
		status,
		current,
		pageSize,
		userID)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	total, err := q.quoteRepo.CountChannelConfigs(
		ctx,
		channelID,
		status,
		userID)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	return result, total, nil
}
