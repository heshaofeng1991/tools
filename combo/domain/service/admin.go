package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	ent "combo/ent"
	admin "combo/gen/admin"
	repository "combo/infrastructure/repository/admin"

	"goa.design/goa/v3/security"
)

// NewAdmin returns the admin service implementation.
func NewAdmin(entClient *ent.Client, logger *log.Logger) admin.Service {
	return &basicAdminService{
		entClient: entClient,
		logger:    logger,
		admin:     repository.New(entClient),
	}
}

// basicAdminService implements Service interface.
type basicAdminService struct {
	entClient *ent.Client
	admin     repository.Admin
	logger    *log.Logger
}

// JWTAuth implements the authorization logic for service "web" for the "jwt"
// security scheme.
func (b basicAdminService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

func (b basicAdminService) GetAnyDoorConfigListEndpoint(ctx context.Context, pagination *admin.Pagination) (res *admin.GetAnyDoorConfigListResp, err error) {
	res, err = b.admin.GetAnyDoorConfigList(ctx, pagination)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) GetAnyDoorConfigInfo(ctx context.Context, req *admin.GetConfigInfoByIDReq) (res *admin.GetAnyDoorConfigInfoResp, err error) {
	res, err = b.admin.GetAnyDoorConfigInfo(ctx, req)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) AddAnyDoorConfig(ctx context.Context, req *admin.AddOrUpdateAnyDoorConfigReq) (res *admin.BaseResponse, err error) {
	res, err = b.admin.AddAnyDoorConfig(ctx, req)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) UpdateAnyDoorConfig(ctx context.Context, req *admin.AddOrUpdateAnyDoorConfigReq) (res *admin.BaseResponse, err error) {
	res, err = b.admin.UpdateAnyDoorConfig(ctx, req)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) AddBoxGoods(ctx context.Context, req *admin.BoxGoodsReq) (res *admin.BaseResponse, err error) {
	res, err = b.admin.AddBoxGoods(ctx, req)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) UpdateBoxGoods(ctx context.Context, req *admin.UpBoxGoodsReq) (res *admin.BaseResponse, err error) {
	res, err = b.admin.UpdateBoxGoods(ctx, req)

	return res, errors.Unwrap(err)
}

func (b basicAdminService) DeleteBoxGoods(ctx context.Context, req *admin.GetConfigInfoByIDReq) (res *admin.BaseResponse, err error) {
	res, err = b.admin.DeleteBoxGoods(ctx, req)

	return res, errors.Unwrap(err)
}
