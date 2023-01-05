/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    create_channel_attribute
	@Date    2022/5/20 12:10
	@Desc
*/

package command

import (
	"context"

	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type CreateChannelAttributeHandler struct {
	quoteRepo domainRepository.ChannelRepository
}

func NewCreateChannelAttributeHandler(quoteRepo domainRepository.ChannelRepository) CreateChannelAttributeHandler {
	return CreateChannelAttributeHandler{
		quoteRepo: quoteRepo,
	}
}

func (q CreateChannelAttributeHandler) Handle(
	ctx context.Context,
	attribute string,
) (int32, error) {
	updateFn := (func(ctx context.Context, attribute string) (int32, error))(nil)

	_, err := q.quoteRepo.CreateChannelAttributes(
		ctx,
		attribute,
		updateFn)
	if err != nil {
		return 0, errors.Wrap(err, "")
	}

	return 1, nil
}
