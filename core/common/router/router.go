package router

import (
	"core/common"
	"core/middlewares/api"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Engine(middleware ...gin.HandlerFunc) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery(), api.Logger())
	engine.Use(middleware...)
	engine.GET(common.HeartbeatPath, Check)
	return engine
}

func Check(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ok")
}
