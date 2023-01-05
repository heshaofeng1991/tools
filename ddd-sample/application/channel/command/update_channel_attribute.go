/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel_attribute
	@Date    2022/5/20 12:10
	@Desc
*/

package command

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelAttributeHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateChannelAttributeHandler(quoteRepo domainRepository.ChannelRepository) UpdateChannelAttributeHandler {
	return UpdateChannelAttributeHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelAttributeHandler) Handle(
	ctx context.Context,
	attribute string,
	channelIDs []int64,
) ([]*domainEntity.UpdateChannelAttribute, error) {
	result, err := q.quoteRepo.UpdateChannelAttributes(
		ctx,
		attribute,
		channelIDs)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
