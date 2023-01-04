package router

import (
	"admin/handler"
	common "core/common/router"
	middle "core/middlewares/api"
	"proto/admin"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	engine := common.Engine(middle.Trace())

	admin.RegisterAdminHTTPHandler(engine.Group("/"), &handler.AdminServer{})

	return engine
}
