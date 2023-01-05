/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    query_channel_cost_by_volume.go
	@Date    2022/5/13 09:43
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetChannelCostInfoByVolumeHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetChannelCostInfoByVolumeHandler(quoteRepo domainRepository.Repository) GetChannelCostInfoByVolumeHandler {
	return GetChannelCostInfoByVolumeHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetChannelCostInfoByVolumeHandler) Handle(ctx context.Context, countryCode string,
	actualGrams int, volume int64,
	volumeChannelCostBatches []*domainEntity.ChannelBatch,
) ([]*domainEntity.ChannelCostRsp, error) {
	result, err := q.quoteRepo.GetChannelCostByVolume(ctx, countryCode, actualGrams, volume, volumeChannelCostBatches)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
