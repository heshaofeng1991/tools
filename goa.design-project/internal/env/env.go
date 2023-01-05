package env

import "os"

func getEnv(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

var OpenAPIDomain = getEnv("OPEN_API_DOMAIN", "http://0.0.0.0:80")

var MysqlDsn = getEnv("MYSQL_DSN", "root:@tcp(127.0.0.1:3306)/test")
