/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    country_zone
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

type GetCountryZoneHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetCountryZoneHandler(quoteRepo domainRepository.Repository) GetCountryZoneHandler {
	return GetCountryZoneHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetCountryZoneHandler) Handle(ctx context.Context,
	zipCode, countryCode string,
) ([]*domainEntity.CountryZone, error) {
	result, err := q.quoteRepo.GetCountryZone(ctx, zipCode, countryCode)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
