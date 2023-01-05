package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"

	"time"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

func (BaseMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").SchemaType(map[string]string{
			dialect.MySQL: "bigint(20)", // Override MySQL.
		}).Comment("主键").Unique(),

		field.Time("created_time").Immutable().Default(time.Now),

		field.Time("updated_time").Default(time.Now).UpdateDefault(time.Now),

		field.Time("deleted_time").SchemaType(map[string]string{
			dialect.MySQL: "timestamp", // Override MySQL.
		}).Optional(),

		field.Int32("revision").SchemaType(map[string]string{
			dialect.MySQL: "int(11)", // Override MySQL.
		}).Optional().Comment("乐观锁"),

		field.String("created_by").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("创建人"),

		field.String("updated_by").SchemaType(map[string]string{
			dialect.MySQL: "varchar(24)", // Override MySQL.
		}).Optional().Comment("更新人"),
	}
}
