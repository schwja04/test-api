package builders

import "fmt"

type PostgresConnectionStringBuilder struct {
	host     string
	user     string
	password string
	database string
	port     int
}

type IPostgresConnectionStringBuilder interface {
	WithHost(host string) IPostgresConnectionStringBuilder
	WithUser(user string) IPostgresConnectionStringBuilder
	WithPassword(password string) IPostgresConnectionStringBuilder
	WithDatabase(database string) IPostgresConnectionStringBuilder
	WithPort(port int) IPostgresConnectionStringBuilder
	Build() string
}

func NewConnectionStringBuilder() *PostgresConnectionStringBuilder {
	return &PostgresConnectionStringBuilder{}
}

func (c *PostgresConnectionStringBuilder) WithHost(host string) *PostgresConnectionStringBuilder {
	c.host = host
	return c
}

func (c *PostgresConnectionStringBuilder) WithUser(user string) *PostgresConnectionStringBuilder {
	c.user = user
	return c
}

func (c *PostgresConnectionStringBuilder) WithPassword(password string) *PostgresConnectionStringBuilder {
	c.password = password
	return c
}

func (c *PostgresConnectionStringBuilder) WithDatabase(database string) *PostgresConnectionStringBuilder {
	c.database = database
	return c
}

func (c *PostgresConnectionStringBuilder) WithPort(port int) *PostgresConnectionStringBuilder {
	c.port = port
	return c
}

func (c *PostgresConnectionStringBuilder) Build() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database)
}
