package environment

import (
	"os"
)

/*
	parseTime=true : []byte -> time.Time
	multiStatements : multiple queries on exec
*/
var dbConfig = "?parseTime=true&multiStatements=true"

type Config struct {
	DbPath           string
	RedisPath        string
	RedisDbs         map[string]int
	DbMaxConnections int
	ApiAddress       string
	JwtKey           []byte
	GwiUser          string
}

func LoadConfig() *Config {
	return &Config{
		DbPath:           getEnv("GWI_MYSQL_PATH", "user:password@tcp(localhost:33066)/gwi"+dbConfig),
		RedisPath:        getEnv("GWI_REDIS_PATH", "localhost:63799"),
		RedisDbs:         map[string]int{"0": 0},
		DbMaxConnections: 50,
		ApiAddress:       getEnv("ADDRESS", ":8080"),
		JwtKey:           []byte("gwi"),
		GwiUser:          "gwiuser",
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
