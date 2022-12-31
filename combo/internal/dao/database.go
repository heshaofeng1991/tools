// Package dao for db.
package dao

import (
	"database/sql"

	ent "combo/ent"
	"combo/internal/env"
	entsql "entgo.io/ent/dialect/sql"
	log "github.com/sirupsen/logrus"

	// import mysql driver.
	_ "github.com/go-sql-driver/mysql"
)

// Open for mysql connect of ent.
func Open() (entClient *ent.Client) {
	mysqlDSN := env.MysqlDsn

	dbConn, err := sql.Open("mysql", mysqlDSN+"?charset=utf8mb4&parseTime=True&loc=UTC")
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Fatalf("failed connecting database")
	}

	// if db.Ping() != nil {
	//	logrus.WithFields(logrus.Fields{
	//		"err": err,
	//	}).Fatalf("failed connecting database")
	//	panic("failed connecting database")
	// }

	dbConn.SetMaxIdleConns(10) //nolint:gomnd

	drv := entsql.OpenDB("mysql", dbConn)

	entClient = ent.NewClient(ent.Driver(drv)).Debug()

	// if err := entClient.Schema.Create(context.Background()); err != nil {
	//	logrus.Fatalf("failed running database migration: %v", err)
	// }

	return entClient
}
