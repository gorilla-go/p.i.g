package Mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Mysql struct {
	DatabasePath string
	Port         int
	DatabaseName string
	UserName     string
	Password     string
	db           *sql.DB
}

func (m *Mysql) Initialize(config MysqlConfig) *Mysql {
	m.UserName = config.UserName
	m.Password = config.Password
	m.DatabaseName = config.DatabaseName
	m.DatabasePath = config.DatabasePath
	m.Port = config.Port

	database, err := sql.Open(
		"mysql",
		fmt.Sprintf(
			"%s:%s@%s:%d/%s",
			m.UserName,
			m.Password,
			m.DatabasePath,
			m.Port,
			m.DatabaseName,
		),
	)

	m.db = database
	if err != nil {
		panic(err)
	}
	database.SetMaxOpenConns(config.MaxOpenConns)
	database.SetConnMaxLifetime(time.Duration(config.MaxConnLifetime) * time.Second)

	database.SetMaxIdleConns(config.MaxIdleConns)
	database.SetConnMaxIdleTime(time.Duration(config.MaxIdleLifetime) * time.Second)
	return m
}

func (m *Mysql) GetRow(sql string, args ...interface{}) {
	err := m.db.QueryRow(sql, args...).Scan()
	if err != nil {
		panic(err)
	}
}
