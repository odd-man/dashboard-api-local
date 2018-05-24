package influxdb

import (
	"github.com/influxdata/influxdb/client/v2"
)

// InfluxDBConnection provide the db service for influxdb
type InfluxDBConnection struct {
	Client *client.Client
}

// Connect get the conn
func (c *InfluxDBConnection) Connect() {
	client, err := client.NewHTTPClient(client.HTTPConfig{
		// Addr:     Addr,
		// Username: username,
		// Password: password,
	})
	if err != nil {

	}
	c.Client = &client
}
