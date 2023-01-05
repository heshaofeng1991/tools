package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CbBoxGoodsConfig holds the schema definition for the CbBoxGoodsConfig entity.
type CbBoxGoodsConfig struct {
	ent.Schema
}

// Mixin of the CbBoxGoodsConfig.
func (CbBoxGoodsConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CbBoxGoodsConfig.
func (CbBoxGoodsConfig) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("type").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Comment("类型 1连击赏赏品 2任意门赏品"),

		field.String("parent_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("所属赏套id"),

		field.String("item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("物品id 真实商品id"),

		field.String("mall_item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("商店商品id 商店商品id"),

		field.String("mall_type").SchemaType(map[string]string{
			dialect.MySQL: "varchar(2)", // Override MySQL.
		}).Optional().Comment("商店商品类型 1普通2预售3秒杀10nft"),

		field.String("nft_item_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("nft商品id nft商品id"),

		field.String("specification").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("sku规格"),

		field.String("show_winning_rate").SchemaType(map[string]string{
			dialect.MySQL: "varchar(32)", // Override MySQL.
		}).Optional().Comment("显示中奖率"),

		field.Int32("item_num").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("物品数量"),

		field.Int32("quality_type").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Optional().Comment("物品品质类型"),

		field.Int32("bottom_num_user").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Default(0).Comment("用户沉底次数"),

		field.Int32("bottom_num_promoter").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Default(0).Comment("主播沉底次数"),
	}
}

// Edges of the CbBoxGoodsConfig.
func (CbBoxGoodsConfig) Edges() []ent.Edge {
	return nil
}
