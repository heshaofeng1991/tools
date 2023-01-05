/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel_recommend
	@Date    2022/5/25 12:38
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelRecommendHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelRecommendHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelRecommendHandler {
	return CreateChannelRecommendHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelRecommendHandler) Handle(
	ctx context.Context,
	countryCode string,
	channelID int64,
) (int64, error) {
	id, err := q.quoteRepo.CreateChannelRecommend(
		ctx,
		countryCode,
		channelID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return id, nil
}
