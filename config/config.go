package config

type ServerConfig struct {
	URL string
}

var Server ServerConfig

func init() {
	Server.URL = "http://192.168.0.65:5050"
}
