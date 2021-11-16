package sql

import (
	"context"
	"database/sql"
	"log"
	"time"

	"gwi/platform2.0-go-challenge/environment"

	_ "github.com/go-sql-driver/mysql"
)

// BasicConnection interface to database connection object.
type BasicConnection interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// BasicConnectionWithTransactions interface.
type BasicConnectionWithTransactions interface {
	BasicConnection
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
}

// Client is a wrap that enhances the standard sql client some transactional utilities.
type Client struct {
	*sql.DB
}

// NewDBClient : Returns a new SQL client based on environment configuration.
func NewDBClient(conf *environment.Config) (*Client, error) {
	conn, err := sql.Open("mysql", conf.DbPath)
	if err != nil {
		log.Printf("error on opening connection with db: %s", err.Error())
		return nil, err
	}

	conn.SetMaxOpenConns(conf.DbMaxConnections)
	conn.SetMaxIdleConns(conf.DbMaxConnections)
	conn.SetConnMaxLifetime(300 * time.Second)

	client := Client{
		DB: conn,
	}

	return &client, nil
}
