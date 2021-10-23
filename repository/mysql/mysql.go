package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/aditya37/file-service/repository"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

type (
	mysql struct {
		db *sql.DB
	}
	MysqlConfig struct {
		Host              string
		Port              int
		Name              string
		User              string
		Password          string
		MaxConnection     int
		MaxIdleConnection int
	}
)

func NewMysqlDataSource(param MysqlConfig) (repository.DBReadWriter, error) {
	connURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		param.User,
		param.Password,
		param.Host,
		param.Port,
		param.Name,
	)

	db, err := sql.Open("mysql", connURL)
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf(
		"MySQL Connection %s:%s@tcp(%s:%d)/%s",
		param.User,
		"********************",
		param.Host,
		param.Port,
		param.Name,
	))

	if param.MaxConnection > 0 {
		db.SetMaxOpenConns(param.MaxConnection)
	}
	if param.MaxIdleConnection > 0 {
		db.SetMaxIdleConns(param.MaxIdleConnection)
	}

	// do migration
	if err := goose.SetDialect("mysql"); err != nil {
		return nil, err
	}
	// Migration db schema for MySQL
	if err := goose.Up(db, "migration"); err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &mysql{
		db: db,
	}, nil
}

func (m *mysql) Close() error {
	return m.Close()
}
