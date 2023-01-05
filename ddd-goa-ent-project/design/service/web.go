package service

import (
	"combo/design/model"

	. "goa.design/goa/v3/dsl"
)

var _ = Service("web", func() {
	Description("The web service performs operations on web")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1/combo_box")
	})

	Method("get_combo_config_list", func() {
		Payload(model.Pagination)
		Result(model.ComboBoxConfigListResp)
		HTTP(func() {
			GET("/list")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_combo_config_info", func() {
		Payload(model.GetConfigInfoByIDReq)
		Result(model.ComboBoxConfigInfoResp)
		HTTP(func() {
			GET("/info")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
