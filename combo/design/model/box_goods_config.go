package model

import (
	. "goa.design/goa/v3/dsl"
)

// BoxGoodsReq 赏箱商品.
var BoxGoodsReq = Type("BoxGoodsReq", func() {
	Field(1, "type", Int, "type 类型 1 连击赏赏品 2 任意门赏品")
	Field(2, "parent_id", String, "parent_id 所属赏套id")
	Field(3, "item_id", String, "item_id 物品id 真实商品id")
	Field(4, "mall_item_id", String, "mall_item_id 商店商品id")
	Field(5, "mall_type", String, "mall_type 商店商品类型 1 普通 2 预售 3 秒杀 10 nft")
	Field(6, "nft_item_id", String, "mall_item_id nft商品id")
	Field(7, "specification", String, "specification sku规格")
	Field(8, "show_winning_rate", String, "show_winning_rate 显示中奖率")
	Field(9, "item_num", Int, "item_num 物品数量")
	Field(10, "quality_type", Int, "quality_type 物品品质类型")
	Field(11, "bottom_num_user", Int, "bottom_num_user 用户沉底次数")
	Field(12, "bottom_num_promoter", Int, "bottom_num_promoter 主播沉底次数")
	Extend(AuthToken)
	Required("type", "parent_id", "item_id", "mall_item_id", "mall_type",
		"nft_item_id", "specification", "show_winning_rate", "item_num",
		"quality_type", "bottom_num_user", "bottom_num_promoter")
})

// UpBoxGoodsReq 赏箱商品.
var UpBoxGoodsReq = Type("UpBoxGoodsReq", func() {
	Field(1, "id", Int, "id 类型 1 连击赏赏品 2 任意门赏品")
	Field(2, "show_winning_rate", String, "show_winning_rate 显示中奖率")
	Field(3, "quality_type", Int, "quality_type 物品品质类型")
	Field(4, "bottom_num_user", Int, "bottom_num_user 用户沉底次数")
	Field(5, "bottom_num_promoter", Int, "bottom_num_promoter 主播沉底次数")
	Extend(AuthToken)
	Required("id", "show_winning_rate")
})
