package web

import (
	"context"

	"combo/ent"
	comboConf "combo/ent/cbcomboboxconfig"
)

type ComboConfig interface {
	GetComboConfigList(ctx context.Context, entClient *ent.Client,
		page, pageSize int64) ([]*ent.CbComboBoxConfig, error)

	GetComboConfigInfoByID(ctx context.Context, entClient *ent.Client, id int64) (*ent.CbComboBoxConfig, error)
}

type Combo struct {
}

func (Combo) GetComboConfigList(ctx context.Context, entClient *ent.Client,
	page, pageSize int) ([]*ent.CbComboBoxConfig, error) {
	return entClient.CbComboBoxConfig.Query().
		Order(ent.Desc(comboConf.FieldID)).Limit(pageSize).
		Offset((page - 1) * pageSize).All(ctx)
}

func (Combo) GetComboConfigInfoByID(ctx context.Context,
	entClient *ent.Client, id int) (*ent.CbComboBoxConfig, error) {
	return entClient.CbComboBoxConfig.Query().Where(comboConf.IDEQ(int64(id))).First(ctx)
}
