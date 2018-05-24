package backend

import (
	"github.com/seeleteam/dashboard-api/backend/influxdb"
	"github.com/seeleteam/dashboard-api/config"
)

var regStruct map[string]Connection

// Connection provide the db service
type Connection interface {
	Connect()
}

func init() {
	regStruct = make(map[string]Connection)
	regStruct["influxdb"] = &influxdb.InfluxDBConnection{}
}

// GetConnection get the real connection from pool
func GetConnection() Connection {
	var conn Connection
	str := config.Config.Database.Backend
	conn, ok := regStruct[str]
	if !ok {
		return nil
	}
	conn.Connect()
	return conn
}
