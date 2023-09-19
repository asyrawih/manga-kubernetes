package dbconn

import (
	"database/sql"
	"fmt"

	"github.com/asyrawih/manga/config"

	_ "github.com/go-sql-driver/mysql"
)

type Option func(db *sql.DB)

func WithOpenMaxConn(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}

func WithMaxIddleConn(n int) Option {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}

// Return new Mysql db instance
func NewMySQLDB(config *config.Config, opts ...Option) (*sql.DB, error) {
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

	// Apply Addtional Max Open Connections
	for _, opt := range opts {
		opt(db)
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
