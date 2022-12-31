package web

import (
	"combo/gen/web"
	"combo/infrastructure/repository"

	"context"
	"errors"
)

var (
	comboBoxConf Combo
)

func (r repo) GetComboConfigList(ctx context.Context, p *web.Pagination) (res *web.ComboBoxConfigListResp, err error) {
	res = &web.ComboBoxConfigListResp{}

	pageSize, page := repository.GetPageSizeAndCurrent(&p.Page, &p.PageSize)

	result, err := comboBoxConf.GetComboConfigList(ctx, r.entClient, page, pageSize)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	for _, val := range result {
		res.List = append(res.List, &web.ComboBoxConfigList{
			AdminUserID: &val.AdminUserID,
			// ObjectID:    val.ObjectID,
			Name:  val.Name,
			Desc:  val.Desc,
			Img:   val.Img,
			ImgBg: val.ImgBg,
			// Index:       val.Index,
			SellPrice: &val.SellPrice,
			// Status:      val.Status,
			BoxType:    nil,
			Revision:   nil,
			ID:         nil,
			CreateTime: nil,
			UpdateTime: nil,
			CreateBy:   nil,
			UpdateBy:   nil,
		})
	}

	total, err := r.entClient.CbComboBoxConfig.Query().Count(ctx)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	newTotal := int64(total)

	res.Total = &newTotal

	return
}

func (r repo) GetComboConfigInfo(ctx context.Context, p *web.GetConfigInfoByIDReq) (res *web.ComboBoxConfigInfoResp, err error) {
	res = &web.ComboBoxConfigInfoResp{}

	result, err := comboBoxConf.GetComboConfigInfoByID(ctx, r.entClient, p.ID)
	if err != nil {
		return nil, errors.Unwrap(err)
	}

	res.Name = result.Name

	return
}
