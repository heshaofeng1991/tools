/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_attribute
	@Date    2022/5/20 12:09
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelAttributeInfoHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewGetChannelAttributeInfoHandler(quoteRepo domainRepository.ChannelRepository) GetChannelAttributeInfoHandler {
	return GetChannelAttributeInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelAttributeInfoHandler) Handle(
	ctx context.Context,
) ([]*domainEntity.ChannelAttribute, error) {
	result, err := q.quoteRepo.GetChannelAttributes(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
