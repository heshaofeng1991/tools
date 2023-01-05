/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    update_channel_cost_status
	@Date    2022/5/25 23:58
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type UpdateChannelCostStatusHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewUpdateUpdateChannelCostStatusHandler(
	quoteRepo domainRepository.ChannelRepository,
) UpdateChannelCostStatusHandler {
	return UpdateChannelCostStatusHandler{
		quoteRepo: quoteRepo,
	}
}

func (q UpdateChannelCostStatusHandler) Handle(
	ctx context.Context,
	ids []int64,
	countryCodes []string,
	status bool,
) (int8, error) {
	_, err := q.quoteRepo.UpdateChannelCostStatus(
		ctx,
		ids,
		countryCodes,
		status)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}
