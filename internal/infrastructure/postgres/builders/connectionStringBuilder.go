package builders

import "fmt"

type ConnectionStringBuilder struct {
	host     string
	user     string
	password string
	database string
	port     int
}

func NewConnectionStringBuilder() *ConnectionStringBuilder {
	return &ConnectionStringBuilder{}
}

func (c *ConnectionStringBuilder) WithHost(host string) *ConnectionStringBuilder {
	c.host = host
	return c
}

func (c *ConnectionStringBuilder) WithUser(user string) *ConnectionStringBuilder {
	c.user = user
	return c
}

func (c *ConnectionStringBuilder) WithPassword(password string) *ConnectionStringBuilder {
	c.password = password
	return c
}

func (c *ConnectionStringBuilder) WithDatabase(database string) *ConnectionStringBuilder {
	c.database = database
	return c
}

func (c *ConnectionStringBuilder) WithPort(port int) *ConnectionStringBuilder {
	c.port = port
	return c
}

func (c *ConnectionStringBuilder) Build() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		c.user,
		c.password,
		c.host,
		c.port,
		c.database)
}
