package model

import (
	. "goa.design/goa/v3/dsl"
)

var GetConfigInfoByIDReq = Type("GetConfigInfoByIDReq", func() {
	Field(1, "id", Int, "id 查询id", func() {
		Minimum(1)
		Example(1)
	})
	Extend(AuthToken)
	Required("id")
})

var ComboBoxConfigInfo = Type("ComboBoxConfigInfo", func() {
	Field(1, "object_id", Int, "_id 兼容mongodb")
	Field(2, "name", String, "name 赏箱名称")
	Field(3, "desc", String, "desc 秒速")
	Field(4, "img", String, "img 赏箱图片")
	Field(5, "img_bg", String, "img_bg 赏箱背景图片")
	Field(6, "index", Int, "index 优先级")
	Field(7, "sell_price", Float64, "sell_price 单发售价")
	Field(8, "status", Int, "status 状态 -1 已删除 0 待上架 1 已上架 2 已下架")
	Field(9, "box_type", String, "box_type 赏箱类型 类型id;分割")

	Required("_id", "name", "desc", "img", "img_bg")
})

var ComboBoxConfigListResp = Type("ComboBoxConfigListResp", func() {
	Field(1, "list", ArrayOf(ComboBoxConfigList), "list 查询数据")
	Field(2, "total", Int64, "总数")
})

var ComboBoxConfigList = Type("ComboBoxConfigList", func() {
	Field(1, "admin_user_id", String, "admin_user_id 负责人(objectid)")
	Extend(ComboBoxConfigInfo)
	Extend(BaseAndRevisionAndAuthor)
})

var ComboBoxConfigInfoResp = Type("ComboBoxConfigInfoResp", func() {
	Field(1, "id", Int, "id 查询id", func() {
		Minimum(1)
		Example(1)
	})
	Field(2, "created_time", Int, "created_time 创建时间")
	Extend(ComboBoxConfigInfo)
})
