/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel
	@Date    2022/5/16 17:03
	@Desc
*/

package command

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateChannelHandler(quoteRepo domainRepository.ChannelRepository) UpdateChannelHandler {
	return UpdateChannelHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelHandler) Handle(
	ctx context.Context,
	params domainEntity.Channel,
) (int, error) {
	status, err := q.quoteRepo.UpdateChannel(ctx,
		params)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return status, nil
}
