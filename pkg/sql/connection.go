package sql

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func connect() (*sqlx.DB, error) {
	conn, err := sqlx.Connect(viper.GetString("db.engine"), viper.GetString("db.connection_url"))
	if err != nil {
		return nil, err
	}

	conn.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	conn.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	conn.SetConnMaxLifetime(0) // 0, connections are reused forever.

	return conn, nil
}
