package dbconn

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/asyrawih/manga/config"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenConns    = 60
	connMaxLifetime = 120
	maxIdleConns    = 30
	connMaxIdleTime = 20
)

// Return new Mysql db instance
func NewMySQLDB(config *config.Config) (*sql.DB, error) {
	dbUser := config.Mysqlusername
	dbPass := config.Mysqlpassword
	dbHost := config.Mysqlhostname
	dbName := config.Mysqldatabase
	dbPort := config.MysqlPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(connMaxLifetime * time.Second)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxIdleTime(connMaxIdleTime * time.Second)
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
