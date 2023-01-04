package handler

import (
	"admin/model"
	"admin/service"
	"context"
	"math/rand"
	"proto/admin"
)

func RandID() int64 {
	return rand.Int63()
}

func (*AdminServer) CreateBrand(ctx context.Context, req *admin.CreateBrandReq) (*admin.CreateBrandRsp, error) {
	rsp := &admin.CreateBrandRsp{
		Data: &admin.CreateBrandData{},
	}
	if err := req.Validate(); err != nil {
		return rsp, err
	}

	id := RandID()

	data := &model.NftIssuer{
		Id:        id,
		Name:      req.Name,
		CoverUrl:  req.CoverUrl,
		SortOrder: int(req.SortOrder),
		Remark:    req.Remark,
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.UpdatedBy,
	}

	if err := service.CreateBrand(ctx, data); err != nil {
		return rsp, err
	}

	rsp.Code = 0
	rsp.Desc = "ok"
	rsp.Data.Id = id

	return rsp, nil
}

func (*AdminServer) UpdateBrand(ctx context.Context, req *admin.UpdateBrandReq) (*admin.CommonRsp, error) {
	rsp := &admin.CommonRsp{}
	if err := req.Validate(); err != nil {
		return rsp, err
	}

	data := &model.NftIssuer{
		Id:         RandID(),
		Name:       req.Name,
		CoverUrl:   req.CoverUrl,
		SortOrder:  int(req.SortOrder),
		Remark:     req.Remark,
		CreatedBy:  req.CreatedBy,
		UpdatedBy:  req.UpdatedBy,
		DeleteFlag: int(req.DeleteFlag),
	}

	if err := service.UpdateBrand(ctx, data); err != nil {
		return rsp, err
	}

	rsp.Code = 0
	rsp.Desc = "ok"

	return rsp, nil
}

func (*AdminServer) GetBrandByID(ctx context.Context, req *admin.GetBrandByIDReq) (*admin.GetBrandByIDRsp, error) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	data := &model.NftIssuer{}

	if err := service.GetBrandByID(ctx, req.BrandId, data); err != nil {
		return nil, err
	}

	rsp := &admin.GetBrandByIDRsp{
		Id:          data.Id,
		Name:        data.Name,
		CoverUrl:    data.CoverUrl,
		SortOrder:   int32(data.SortOrder),
		CreatedBy:   data.CreatedBy,
		UpdatedBy:   data.UpdatedBy,
		Remark:      data.Remark,
		DeleteFlag:  int32(data.DeleteFlag),
		CreatedTime: &data.CreatedAt,
		UpdatedTime: &data.UpdatedAt,
	}

	return rsp, nil
}

func (*AdminServer) MultiGetBrand(ctx context.Context, req *admin.MultiGetBrandReq) (*admin.MultiGetBrandRsp, error) {
	rsp := &admin.MultiGetBrandRsp{
		Data: &admin.MultiGetBrandData{},
	}

	if err := req.Validate(); err != nil {
		return rsp, err
	}

	list, total, err := service.GetBrandList(ctx, req)
	if err != nil {
		return rsp, err
	}

	rsp.Code = 0
	rsp.Desc = "ok"
	rsp.Data.Total = total

	for _, val := range list {
		content := &admin.MultiGetBrandContent{
			Id:          val.Id,
			Name:        val.Name,
			CoverUrl:    val.CoverUrl,
			SortOrder:   int32(val.SortOrder),
			CreatedBy:   val.CreatedBy,
			UpdatedBy:   val.UpdatedBy,
			Remark:      val.Remark,
			DeleteFlag:  int32(val.DeleteFlag),
			CreatedTime: &val.CreatedAt,
			UpdatedTime: &val.UpdatedAt,
		}

		rsp.Data.List = append(rsp.Data.List, content)
	}

	return rsp, nil
}

func (*AdminServer) RemoveBrand(ctx context.Context, req *admin.RemoveBrandReq) (*admin.CommonRsp, error) {
	rsp := &admin.CommonRsp{}
	if err := req.Validate(); err != nil {
		return rsp, err
	}

	data := &model.NftIssuer{
		Id: req.BrandId,
	}

	if err := service.RemoveBrand(ctx, data); err != nil {
		return rsp, err
	}

	rsp.Code = 0
	rsp.Desc = "ok"

	return rsp, nil
}
