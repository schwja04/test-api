package factories

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IConnectionFactory interface {
	GetConnection() (*pgxpool.Conn, error)
	Close()
}

type PostgresConnectionFactory struct {
	connectionString string
	pool             *pgxpool.Pool
}

func NewPostgresConnectionFactory(connectionString string) IConnectionFactory {
	pool, err := pgxpool.New(context.Background(), connectionString)
	if err != nil {
		log.Fatal(err)
	}

	return &PostgresConnectionFactory{
		connectionString: connectionString,
		pool:             pool,
	}
}

func (cf *PostgresConnectionFactory) GetConnection() (*pgxpool.Conn, error) {
	conn, err := cf.pool.Acquire(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (cf *PostgresConnectionFactory) Close() {
	cf.pool.Close()
}
