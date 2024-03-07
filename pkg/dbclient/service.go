package dbclient

type pgClient struct {
	masterDBC DB
}

// New initializes a new user repository using the provided database configuration.
func New(dbc DB) (Client, error) {
	return &pgClient{
		masterDBC: dbc,
	}, nil
}

// DB returns the database client
func (c *pgClient) DB() DB {
	return c.masterDBC
}

// Close closes the database client
func (c *pgClient) Close() error {
	if c.masterDBC != nil {
		c.masterDBC.Close()
	}

	return nil
}
