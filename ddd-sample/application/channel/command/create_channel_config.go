/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel_config
	@Date    2022/5/25 12:39
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelConfigHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelConfigHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelConfigHandler {
	return CreateChannelConfigHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelConfigHandler) Handle(
	ctx context.Context,
	countryCodes []string,
	ids []int64,
	userID int64,
) (int64, error) {
	id, err := q.quoteRepo.CreateChannelConfig(
		ctx,
		countryCodes,
		ids,
		userID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return id, nil
}
