package web

import (
	"context"

	"combo/gen/web"
)

type Web interface {
	GetComboConfigList(ctx context.Context, p *web.Pagination) (res *web.ComboBoxConfigListResp, err error)
	GetComboConfigInfo(ctx context.Context, p *web.GetConfigInfoByIDReq) (res *web.ComboBoxConfigInfoResp, err error)
}
