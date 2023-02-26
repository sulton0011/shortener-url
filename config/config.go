package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
	ErrorModel  = "!!!Error"
	InfoModel  = "---Info"
	ErrorStyle  = "-->"

	// service name
	SvcsAnalytic = "AnalyticService"

	PostgresCtxTimeout = time.Second * 30 // 30s
	PostgresLogTime    = time.Second * 5  // 5s
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string

	HTTPPort   string
	HTTPScheme string

	PostgresHost           string
	PostgresPort           int
	PostgresUser           string
	PostgresPassword       string
	PostgresDatabase       string
	PostgresMaxConnections int32

	SecretKey string

	DefaultPage  string
	DefaultLimit string

	RedisHost     string
	RedisPort     int
	RedisPassword string
}

// Load ...
func Load() Config {

	envFileName := cast.ToString(getOrReturnDefaultValue("ENV_FILE_PATH", "./.env"))
fmt.Println(envFileName)
	if err := godotenv.Load(envFileName); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "shortener_url"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))
	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":90"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 32768))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "husniddinnurmamatov"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", ""))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "admin"))
	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	config.SecretKey = cast.ToString(getOrReturnDefaultValue("SECRET_KEY", "aLlhFUx7onmV2IBv6"))

	config.DefaultPage = cast.ToString(getOrReturnDefaultValue("DEFAULT_PAGE", "1"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "127.0.0.1"))
	config.RedisPort = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))
	config.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", ""))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
