package service

import (
	"context"
	"errors"
	"fmt"
	"log"

	ent "combo/ent"
	web "combo/gen/web"
	repository "combo/infrastructure/repository/web"

	"goa.design/goa/v3/security"
)

// NewWeb returns the web service implementation.
func NewWeb(entClient *ent.Client, logger *log.Logger) web.Service {
	return &basicWebService{
		entClient: entClient,
		logger:    logger,
		web:       repository.New(entClient),
	}
}

// basicWebService implements Service interface.
type basicWebService struct {
	entClient *ent.Client
	web       repository.Web
	logger    *log.Logger
}

// JWTAuth implements the authorization logic for service "web" for the "jwt"
// security scheme.
func (b basicWebService) JWTAuth(ctx context.Context, token string, scheme *security.JWTScheme) (context.Context, error) {
	return ctx, fmt.Errorf("not implemented")
}

func (b basicWebService) GetComboConfigList(ctx context.Context, pagination *web.Pagination) (res *web.ComboBoxConfigListResp, err error) {
	res, err = b.web.GetComboConfigList(ctx, pagination)

	return res, errors.Unwrap(err)
}

func (b basicWebService) GetComboConfigInfo(ctx context.Context, req *web.GetConfigInfoByIDReq) (res *web.ComboBoxConfigInfoResp, err error) {
	res, err = b.web.GetComboConfigInfo(ctx, req)

	return res, errors.Unwrap(err)
}
