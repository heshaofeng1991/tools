package service

import (
	"airmart-core/types"
	"context"

	"admin/global"
	"admin/model"
	"airmart-core/common"
	pb "proto/admin"
)

func CreateBrand(ctx context.Context, data *model.NftIssuer) error {
	return global.Srv.DB.WithContext(ctx).Create(data).Error
}

func UpdateBrand(ctx context.Context, data *model.NftIssuer) error {
	return global.Srv.DB.WithContext(ctx).Updates(data).Error
}

func GetBrandByID(ctx context.Context, id int64, result interface{}) (err error) {
	return global.Srv.DB.WithContext(ctx).Model(
		&model.NftIssuer{}).Where("id = ?", id).Find(result).Error
}

func GetBrandList(ctx context.Context, req *pb.MultiGetBrandReq) (list []*model.NftIssuer, total int64, err error) {
	//根据实标查询条件写
	where := map[string]interface{}{}

	if req.Name != "" {
		where["name"] = req.Name
	}

	timeWhere := ""
	if req.StartTime.String() != "" && req.EndTime.String() != "" {
		timeWhere = "created_time >= ? AND created_time <= ?"
	}

	total, err = GetBrandCount(ctx, where)
	if err != nil {
		return
	}

	pages := types.Pages{
		Page:     int(req.Page),
		PageSize: int(req.PageSize),
	}

	res := global.Srv.DB.WithContext(ctx).Model(model.NftIssuer{}).Scopes(common.Paginate(pages)).Order("created_time desc")
	if timeWhere != "" {
		err = res.Find(&list, where).Where(timeWhere, req.EndTime, req.EndTime).Error
	} else {
		err = res.Find(&list, where).Error
	}

	return list, total, err
}

func GetBrandCount(ctx context.Context, where map[string]interface{}) (int64, error) {
	var total int64
	err := global.Srv.DB.WithContext(ctx).Model(
		&model.NftIssuer{}).Where(where).Count(&total).Error
	return total, err
}

func RemoveBrand(ctx context.Context, data *model.NftIssuer) error {
	return global.Srv.DB.WithContext(ctx).Delete(data, data.Id).Error
}
