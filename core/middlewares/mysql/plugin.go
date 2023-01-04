package mysql

import (
	"context"
	"core/common/log"

	"go.opentelemetry.io/otel"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type GormTracing struct {
	ServiceName string
}

const (
	spankey = "gorm-tracing"

	// 自定义事件名称
	_eventBeforeCreate = "gorm-tracing-event:before_create"
	_eventAfterCreate  = "gorm-tracing-event:after_create"
	_eventBeforeUpdate = "gorm-tracing-event:before_update"
	_eventAfterUpdate  = "gorm-tracing-event:after_update"
	_eventBeforeQuery  = "gorm-tracing-event:before_query"
	_eventAfterQuery   = "gorm-tracing-event:after_query"
	_eventBeforeDelete = "gorm-tracing-event:before_delete"
	_eventAfterDelete  = "gorm-tracing-event:after_delete"
	_eventBeforeRow    = "gorm-tracing-event:before_row"
	_eventAfterRow     = "gorm-tracing-event:after_row"
	_eventBeforeRaw    = "gorm-tracing-event:before_raw"
	_eventAfterRaw     = "gorm-tracing-event:after_raw"

	// 自定义 span 的操作名称
	_opCreate = "create"
	_opUpdate = "update"
	_opQuery  = "query"
	_opDelete = "delete"
	_opRow    = "row"
	_opRaw    = "raw"
)

func (*GormTracing) Name() string {
	return "GormTracing"
}

func (*GormTracing) Initialize(db *gorm.DB) (err error) {
	for _, e := range []error{
		db.Callback().Create().Before("gorm:create").Register(_eventBeforeCreate, beforeCreate),
		db.Callback().Create().After("gorm:create").Register(_eventAfterCreate, after),
		db.Callback().Update().Before("gorm:update").Register(_eventBeforeUpdate, beforeUpdate),
		db.Callback().Update().After("gorm:update").Register(_eventAfterUpdate, after),
		db.Callback().Query().Before("gorm:query").Register(_eventBeforeQuery, beforeQuery),
		db.Callback().Query().After("gorm:query").Register(_eventAfterQuery, after),
		db.Callback().Delete().Before("gorm:delete").Register(_eventBeforeDelete, beforeDelete),
		db.Callback().Delete().After("gorm:delete").Register(_eventAfterDelete, after),
		db.Callback().Row().Before("gorm:row").Register(_eventBeforeRow, beforeRow),
		db.Callback().Row().After("gorm:row").Register(_eventAfterRow, after),
		db.Callback().Raw().Before("gorm:raw").Register(_eventBeforeRaw, beforeRaw),
		db.Callback().Raw().After("gorm:raw").Register(_eventAfterRaw, after),
	} {
		if e != nil {
			return e
		}
	}
	return
}

func _injectBefore(db *gorm.DB, op string) {
	if db == nil {
		return
	}

	if db.Statement == nil || db.Statement.Context == nil {
		db.Logger.Error(context.TODO(), "未定义 db.Statement 或 db.Statement.Context")
		return
	}
	op = "mysql-" + op
	spanName, ok := db.InstanceGet("spanName")
	if !ok {
		spanName = op
	}

	_, span := otel.Tracer(op).Start(db.Statement.Context, spanName.(string))
	db.InstanceSet(spankey, span)
}

func after(db *gorm.DB) {
	if db.Statement == nil || db.Statement.Context == nil {
		return
	}
	span, ok := db.InstanceGet(spankey)
	if !ok {
		return
	}
	traceSpan := span.(trace.Span)
	sql := db.Dialector.Explain(db.Statement.SQL.String(), db.Statement.Vars...)

	logger := log.New(db.Statement.Context)
	traceSpan.SetAttributes(
		semconv.DBStatementKey.String(sql),
	)
	if db.Statement.Error != nil {
		logger.Error(db.Statement.Error)
		traceSpan.RecordError(db.Statement.Error, trace.WithStackTrace(true))
	}
	logger.Infof("sql:%s", sql)
	traceSpan.End()
}

func beforeCreate(db *gorm.DB) {
	_injectBefore(db, _opCreate)
}

func beforeUpdate(db *gorm.DB) {
	_injectBefore(db, _opUpdate)
}

func beforeQuery(db *gorm.DB) {
	_injectBefore(db, _opQuery)
}

func beforeDelete(db *gorm.DB) {
	_injectBefore(db, _opDelete)
}

func beforeRow(db *gorm.DB) {
	_injectBefore(db, _opRow)
}

func beforeRaw(db *gorm.DB) {
	_injectBefore(db, _opRaw)
}
