package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CbUserOpenBoxOrder holds the schema definition for the CbUserOpenBoxOrder entity.
type CbUserOpenBoxOrder struct {
	ent.Schema
}

// Mixin of the CbUserOpenBoxOrder.
func (CbUserOpenBoxOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CbUserOpenBoxOrder.
func (CbUserOpenBoxOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("object_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("mongodb Object Id"),

		field.String("type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("类型"),

		field.String("combo_box_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("连击赏id"),

		field.String("door_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("任意门id"),

		field.Int32("combo").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("连击次数"),

		field.String("user_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("用户类型"),

		field.String("user_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("用户id"),

		field.String("nickname").SchemaType(map[string]string{
			dialect.MySQL: "varchar(100)", // Override MySQL.
		}).Optional().Comment("用户昵称"),

		field.String("user_phone").SchemaType(map[string]string{
			dialect.MySQL: "varchar(11)", // Override MySQL.
		}).Optional().Comment("用户手机号"),

		field.String("status").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("订单状态"),

		field.Float("purchase_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("单开价格"),

		field.Float("pay_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("总支付金额"),

		field.Time("pay_time").SchemaType(map[string]string{
			dialect.MySQL: "datetime", // Override MySQL.
		}).Optional().Comment("支付时间"),

		field.String("discount_user_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("优惠券id(预留)"),

		field.String("recharge_order_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("绑定的充值订单"),

		field.Float("balance_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("余额池支付金额"),

		field.Float("token_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("元气石抵扣金额"),

		field.Float("coin_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("元气金币抵扣金额"),

		field.Float("discount_amount").SchemaType(map[string]string{
			dialect.MySQL: "decimal(24,6)", // Override MySQL.
		}).Optional().Comment("优惠券抵扣金额(预留)"),

		field.String("app_channel").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Optional().Comment("设备渠道"),

		field.String("app_version").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Optional().Comment("设备版本号"),

		field.String("ip").SchemaType(map[string]string{
			dialect.MySQL: "varchar(20)", // Override MySQL.
		}).Optional().Comment("ip"),

		field.String("login_ip_address").SchemaType(map[string]string{
			dialect.MySQL: "varchar(255)", // Override MySQL.
		}).Optional().Comment("ip地理位置"),
	}
}

// Edges of the CbUserOpenBoxOrder.
func (CbUserOpenBoxOrder) Edges() []ent.Edge {
	return nil
}
