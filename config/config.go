package config

type _Config struct {
	Database Database
}

// Config config for database
var Config _Config

// Database database scheme
type Database struct {
	Backend string `json:"Backend"`
	Host    string `json:"Host"`
	Port    string `json:"Port"`
	Name    string `json:"Name"`
}
