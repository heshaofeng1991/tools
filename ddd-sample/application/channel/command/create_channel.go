/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel
	@Date    2022/5/16 15:37
	@Desc
*/

package command

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelHandler {
	return CreateChannelHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelHandler) Handle(
	ctx context.Context,
	params domainEntity.Channel,
) (int64, error) {
	id, err := q.quoteRepo.CreateChannel(
		ctx,
		params)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return id, nil
}
