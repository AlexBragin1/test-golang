package db

import (
	"fmt"
	"test/config"
	"time"

	"github.com/jmoiron/sqlx"
)

func GetDB() *sqlx.DB {
	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Global.DB.User, config.Global.DB.Pass,
		config.Global.DB.Host, config.Global.DB.Port,
		config.Global.DB.Name,
	)

	client, err := sqlx.Open("postgres", dataSource)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
