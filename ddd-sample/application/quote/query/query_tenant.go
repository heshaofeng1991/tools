/*
	@Author  johnny
	@Author  johnny.he@nextsmartship.com
	@Version v1.0.0
	@File    tenant
	@Date    2022/5/11 18:02
	@Desc
*/

package query

import (
	"context"

	domainEntity "github.com/heshaofeng1991/ddd-sample/domain/entity"
	domainRepository "github.com/heshaofeng1991/ddd-sample/domain/repository"
	"github.com/pkg/errors"
)

type GetTenantInfoHandler struct {
	quoteRepo domainRepository.Repository
}

func NewGetTenantInfoHandler(quoteRepo domainRepository.Repository) GetTenantInfoHandler {
	return GetTenantInfoHandler{
		quoteRepo: quoteRepo,
	}
}

func (q GetTenantInfoHandler) Handle(ctx context.Context, userID int64) (*domainEntity.Tenant, error) {
	result, err := q.quoteRepo.GetTenantInfo(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return result, nil
}
