package model

import (
	. "goa.design/goa/v3/dsl"
)

var GetAnyDoorConfigList = Type("GetAnyDoorConfigList", func() {
	Field(1, "combo_id", Int, "combo_id 连击赏id")
	Field(2, "combo_bun", Int, "combo_bun 门槛(连击次数)")
	Field(3, "admin_user_id", Int, "admin_user_id 负责人(objectid)", func() {
		Minimum(1)
		Example(1)
	})
	Extend(ComboBoxConfigInfo)
	Extend(BaseAndRevisionAndAuthor)
})

var GetAnyDoorConfigListResp = Type("GetAnyDoorConfigListResp", func() {
	Field(1, "list", ArrayOf(GetAnyDoorConfigList), "list")
	Field(2, "total", Int64, "总数")
})

var GetAnyDoorConfigInfoResp = Type("GetAnyDoorConfigInfoResp", func() {
	Extend(GetAnyDoorConfigList)
})

var AddOrUpdateAnyDoorConfigReq = Type("AddOrUpdateAnyDoorConfigReq", func() {
	Field(1, "id", UInt, "id")
	Field(2, "combo_id", Int, "combo_id 连击赏id")
	Field(3, "combo_bun", Int, "combo_bun 门槛(连击次数)")
	Field(4, "admin_user_id", String, "admin_user_id //负责人(objectid)")
	Field(5, "create_by", String, "create_by 创建人")
	Field(6, "update_by", String, "update_by 更新人")
	Field(7, "revision", Int, "revision 乐观锁")
	Extend(ComboBoxConfigInfo)
	Extend(AuthToken)
})
