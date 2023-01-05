package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CbAnyDoorConfig holds the schema definition for the CbAnyDoorConfig entity.
type CbAnyDoorConfig struct {
	ent.Schema
}

// Mixin of the CbAnyDoorConfig.
func (CbAnyDoorConfig) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the CbAnyDoorConfig.
func (CbAnyDoorConfig) Fields() []ent.Field {
	return []ent.Field{
		field.String("object_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("object id"),

		field.String("name").SchemaType(map[string]string{
			dialect.MySQL: "varchar(80)", // Override MySQL.
		}).Optional().Comment("任意门名称"),

		field.String("desc").SchemaType(map[string]string{
			dialect.MySQL: "varchar(200)", // Override MySQL.
		}).Optional().Comment("描述"),

		field.String("img").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("任意门图片"),

		field.String("img_bg").SchemaType(map[string]string{
			dialect.MySQL: "varchar(1024)", // Override MySQL.
		}).Optional().Comment("任意门背景图片"),

		field.Int32("index").SchemaType(map[string]string{
			dialect.MySQL: "int(8)", // Override MySQL.
		}).Optional().Comment("顺序"),

		field.String("combo_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("连击赏id"),

		field.Int32("status").SchemaType(map[string]string{
			dialect.MySQL: "int(2)", // Override MySQL.
		}).Optional().Comment("状态(预留) -1删除0新增1绑定2解绑"),

		field.Int32("combo_num").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("门槛(连击次数)"),

		field.Float("sell_price").SchemaType(map[string]string{
			dialect.MySQL: "decimal(32,8)", // Override MySQL.
		}).Optional().Comment("单发售价"),

		field.String("admin_user_id").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("负责人(objectid)"),

		field.Text("box_type").SchemaType(map[string]string{
			dialect.MySQL: "text", // Override MySQL.
		}).Optional().Comment("赏箱类型 类型id;分割"),
	}
}

// Edges of the CbAnyDoorConfig.
func (CbAnyDoorConfig) Edges() []ent.Edge {
	return nil
}
