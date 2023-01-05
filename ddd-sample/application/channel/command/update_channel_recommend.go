/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel_recommend
	@Date    2022/5/25 12:38
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelRecommendHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateChannelRecommendHandler(quoteRepo domainRepository.ChannelRepository) UpdateChannelRecommendHandler {
	return UpdateChannelRecommendHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelRecommendHandler) Handle(
	ctx context.Context,
	countryCode string,
	channelID int64,
	isRecommended,
	status bool,
) (int, error) {
	id, err := q.quoteRepo.UpdateChannelRecommend(
		ctx,
		countryCode,
		channelID,
		isRecommended,
		status)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return id, nil
}
