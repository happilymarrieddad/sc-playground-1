package postgres

import (
	"api/types"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"xorm.io/xorm"
)

func NewDB(
	dbUser, dbPass, dbHost, dbPort, dbDatabase string,
) (*xorm.Engine, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?connect_timeout=180&sslmode=disable",
		dbUser, dbPass, dbHost, dbPort, dbDatabase,
	)

	db, err := xorm.NewEngine("pgx", connStr)
	if err != nil {
		return nil, err
	}

	if err = db.Sync(
		&types.Customer{},
		&types.User{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
