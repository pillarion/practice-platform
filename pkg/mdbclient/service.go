package mdbclient

type mdbClient struct {
	cli Cache
}

// NewMdbClient initializes a new DB driver using the provided database configuration.
func NewMdbClient(cli Cache) *mdbClient {
	return &mdbClient{
		cli: cli,
	}
}
