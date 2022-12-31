package service

import (
	"combo/design/model"
	. "goa.design/goa/v3/dsl"
)

var _ = Service("admin", func() {
	Description("The admin service performs operations on admin")

	Security(JWTAuth)

	Error("Unauthorized")

	HTTP(func() {
		Path("/v1")
	})

	Method("get_any_door_config_list", func() {
		Payload(model.Pagination)
		Result(model.GetAnyDoorConfigListResp)
		HTTP(func() {
			GET("/any_door")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("get_any_door_config_info", func() {
		Payload(model.GetConfigInfoByIDReq)
		Result(model.GetAnyDoorConfigInfoResp)
		HTTP(func() {
			GET("/any_door/info")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("add_any_door_config", func() {
		Payload(model.AddOrUpdateAnyDoorConfigReq)
		Result(model.BaseResponse)
		HTTP(func() {
			POST("/any_door/add")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_any_door_config", func() {
		Payload(model.AddOrUpdateAnyDoorConfigReq)
		Result(model.BaseResponse)
		HTTP(func() {
			POST("/any_door/update")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("add_box_goods", func() {
		Payload(model.BoxGoodsReq)
		Result(model.BaseResponse)
		HTTP(func() {
			POST("/box_goods/add")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("update_box_goods", func() {
		Payload(model.UpBoxGoodsReq)
		Result(model.BaseResponse)
		HTTP(func() {
			POST("/box_goods/update")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})

	Method("delete_box_goods", func() {
		Payload(model.GetConfigInfoByIDReq)
		Result(model.BaseResponse)
		HTTP(func() {
			GET("/box_goods/delete")
			Header("Authorization:Authorization")
			Response(StatusOK)
			Response("Unauthorized", StatusUnauthorized)
		})
	})
})
