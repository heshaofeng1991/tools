/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    vat
	@Date    2022/5/11 18:01
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetVatInfoHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetVatInfoHandler(quoteRepo domainRepository.Repository) GetVatInfoHandler {
	return GetVatInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetVatInfoHandler) Handle(ctx context.Context, countryCode string) (*domainEntity.Vat, error) {
	result, err := q.quoteRepo.GetVatInfoByCountryCode(ctx, countryCode)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
