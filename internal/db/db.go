package db

import (
	"fmt"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"
)

type BaseDB struct {
	DB  *sqlx.DB
	Log *zap.SugaredLogger
}

type DBConfig struct {
	Username string
	Password string
	Hostname string
	Port     int
	DBname   string
	Log      *zap.SugaredLogger
}

func NewDB(c DBConfig) (*BaseDB, error) {
	db, err := sqlx.Connect("pgx", fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		c.Username, c.Password, c.Hostname, c.Port, c.DBname))
	if err != nil {
		c.Log.Errorw("DB - NewDB - Error while connecting to DB", "error", err)
		return nil, err
	}

	return &BaseDB{
		DB:  db,
		Log: c.Log,
	}, nil
}
