/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_recommend
	@Date    2022/5/25 22:43
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
	quoteRepo domainRepository.Repository
}

func NewGetChannelRecommendHandler(quoteRepo domainRepository.Repository) GetChannelRecommendHandler {
	return GetChannelRecommendHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelRecommendHandler) Handle(
	ctx context.Context,
	ids []int64,
	countryCode string,
) (*domainEntity.ChannelRecommend, error) {
	result, err := q.quoteRepo.
		GetChannelRecommendsByCondition(ctx,
			ids,
			countryCode)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
