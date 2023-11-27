package postgres

import (
	"database/sql"
	"first_project/config"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Postrgress struct {
	*bun.DB
}

func ConnectToPostgres(pgConf config.Postgres) (*gorm.DB, error) {
	dns := pgConf.GetDNS()

	dbConn, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return dbConn, nil
}

func NewConnectToPostgress(pgConf config.Postgres) (*Postrgress, error) {
	psql := &Postrgress{}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", pgConf.User, pgConf.Password, pgConf.Host, pgConf.Port, pgConf.DbName, pgConf.SslMode)
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	psql.DB = bun.NewDB(sqldb, pgdialect.New())

	return psql, nil
}
