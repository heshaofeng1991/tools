package admin

import (
	"context"

	"combo/gen/admin"
)

type Admin interface {
	GetAnyDoorConfigList(ctx context.Context, pagination *admin.Pagination) (res *admin.GetAnyDoorConfigListResp, err error)
	GetAnyDoorConfigInfo(ctx context.Context, req *admin.GetConfigInfoByIDReq) (res *admin.GetAnyDoorConfigInfoResp, err error)
	AddAnyDoorConfig(ctx context.Context, req *admin.AddOrUpdateAnyDoorConfigReq) (res *admin.BaseResponse, err error)
	UpdateAnyDoorConfig(ctx context.Context, req *admin.AddOrUpdateAnyDoorConfigReq) (res *admin.BaseResponse, err error)

	AddBoxGoods(ctx context.Context, req *admin.BoxGoodsReq) (res *admin.BaseResponse, err error)
	UpdateBoxGoods(ctx context.Context, req *admin.UpBoxGoodsReq) (res *admin.BaseResponse, err error)
	DeleteBoxGoods(ctx context.Context, req *admin.GetConfigInfoByIDReq) (res *admin.BaseResponse, err error)
}
