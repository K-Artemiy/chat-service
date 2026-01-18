package config

import "os"

type Config struct {
	Addr string
	DSN  string
}

func Load() Config {
	addr := os.Getenv("HTTP_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	dsn := os.Getenv("DB_DSN")

	return Config{
		Addr: addr,
		DSN:  dsn,
	}
}
