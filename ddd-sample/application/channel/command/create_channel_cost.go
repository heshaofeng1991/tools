/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel_cost
	@Date    2022/5/17 14:33
	@Desc
*/

package command

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelCostHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelCostHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelCostHandler {
	return CreateChannelCostHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelCostHandler) Handle(
	ctx context.Context,
	costs []*domainEntity.ChannelCost,
) error {
	if err := q.quoteRepo.CreateChannelCost(ctx, costs); err != nil {
		return errors.Wrap(err, "")
	}

	return nil
}
