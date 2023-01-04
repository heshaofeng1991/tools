package service

import (
	"context"

	"gorm.io/gorm"

	"admin/global"
)

type Service struct {
	parentCtx context.Context
	db        *gorm.DB
	tx        *gorm.DB
}

func initService(ctx context.Context, db ...*gorm.DB) *Service {
	dbCtx := global.Srv.DB.WithContext(ctx)

	if len(db) > 0 {
		dbCtx = db[0].WithContext(ctx)
	}

	return &Service{
		parentCtx: ctx,
		db:        dbCtx,
	}
}

type admin struct {
	Service
}

func NewAdmin(ctx context.Context, db ...*gorm.DB) *admin {
	return &admin{*initService(ctx, db...)}
}
