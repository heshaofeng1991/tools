package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CbComboBoxConfig holds the schema definition for the CbComboBoxConfig entity.
type CbComboBoxConfig struct {
	ent.Schema
}

// Mixin of the CbComboBoxConfig.
func (CbComboBoxConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CbComboBoxConfig.
func (CbComboBoxConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("object_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("object id,兼容mongodb"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(80)", // Override MySQL.
		}).Optional().Comment("赏箱名称"),

		field.String("desc").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Optional().Comment("描述"),

		field.String("img").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("赏箱图片"),

		field.String("img_bg").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("赏箱背景图片"),

		field.Int32("index").SchemaType(map[string]string{
			dialect.MySQL: "int(8)", // Override MySQL.
		}).Optional().Comment("优先级"),

		field.Float("sell_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(32,8)", // Override MySQL.
		}).Optional().Comment("单发售价"),

		field.String("admin_user_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("负责人(objectid)"),

		field.Int32("status").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Optional().Comment("状态 -1已删除0待上架1已上架2已下架"),

		field.Text("box_type").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Optional().Comment("赏箱类型 类型id;分割"),
	}
}

// Edges of the CbComboBoxConfig.
func (CbComboBoxConfig) Edges() []ent.Edge {
	return nil
}
