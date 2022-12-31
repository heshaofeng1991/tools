package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CbUserBoxSubOrder holds the schema definition for the CbUserBoxSubOrder entity.
type CbUserBoxSubOrder struct {
	ent.Schema
}

// Mixin of the CbUserBoxSubOrder.
func (CbUserBoxSubOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CbUserBoxSubOrder.
func (CbUserBoxSubOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("object_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("mongodb Object Id"),

		field.String("order_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("主订单id"),

		field.String("item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("物品id"),

		field.String("user_item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("用户背包id"),

		field.String("quality_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("物品品质类型"),

		field.String("mall_item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("商店商品id"),

		field.String("mall_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("商店商品类型"),

		field.String("nft_item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("nft商品id"),

		field.String("specification").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("sku规格"),

		field.Float("pay_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("单个实付价格"),

		field.Float("balance_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("余额支付金额"),

		field.Float("token_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("元气石抵扣金额"),

		field.Float("coin_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("元气金币抵扣金额"),

		field.Float("discount_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("优惠券抵扣金额(预留)"),
	}
}

// Edges of the CbUserBoxSubOrder.
func (CbUserBoxSubOrder) Edges() []ent.Edge {
	return nil
}
