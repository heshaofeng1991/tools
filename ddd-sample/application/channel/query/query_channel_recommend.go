/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_recommend
	@Date    2022/5/25 12:40
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelRecommendHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelRecommendHandler(quoteRepo domainRepository.ChannelRepository) GetChannelRecommendHandler {
	return GetChannelRecommendHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelRecommendHandler) Handle(
	ctx context.Context,
	countryCode, sorter *string,
	channelID *int64,
	status, isRecommended *bool,
	current, pageSize int,
) ([]*domainEntity.ChannelRecommend, int64, error) {
	result, err := q.quoteRepo.GetChannelRecommends(
		ctx,
		countryCode,
		sorter,
		channelID,
		status,
		isRecommended,
		current,
		pageSize)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	total, err := q.quoteRepo.CountChannelRecommends(
		ctx,
		countryCode,
		channelID,
		status,
		isRecommended)
	if err != nil {
		return nil, 0, errors.Wrap(err, "")
	}

	return result, total, nil
}
