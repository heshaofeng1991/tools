/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel_config
	@Date    2022/5/25 12:39
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelConfigHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateChannelConfigHandler(quoteRepo domainRepository.ChannelRepository) UpdateChannelConfigHandler {
	return UpdateChannelConfigHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelConfigHandler) Handle(
	ctx context.Context,
	countryCodes []string,
	ids []int64,
	status bool,
	userID int64,
) (int, error) {
	id, err := q.quoteRepo.UpdateChannelConfig(
		ctx,
		countryCodes,
		ids,
		status,
		userID)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return id, nil
}
