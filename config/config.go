package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	// main
	Environment string
	LogLevel    string
	HttpPort    string

	// services
	UserServiceHost    string
	UserServicePort    int
	ContentServiceHost string
	ContentServicePort int
	AnalyticServiceUrl string

	// redis
	RedisHost             string
	RedisPassword         string
	RedisPort             int
	RedisTimeoutInSeconds int

	// minio
	BucketName          string
	FilesBucketName     string
	JackBucketName      string
	MinioDomain         string
	MinioAccessKeyID    string
	MinioSecretAccesKey string
	MinioLocation       string

	// JWT
	AccessTokenExpireDuration  uint
	RefreshTokenExpireDuration uint

	// keys
	LoginSecretAccessKey  string
	LoginSecretRefreshKey string
}

func Load() Config {
	cfg := Config{}
	// main
	cfg.HttpPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8100"))
	cfg.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", "develop"))
	cfg.LogLevel = cast.ToString(getOrReturnDefaultValue("LOG_LEVEL", "debug"))

	// services
	cfg.UserServiceHost = cast.ToString(getOrReturnDefaultValue("USER_SERVICE_HOST", "localhost"))
	cfg.UserServicePort = cast.ToInt(getOrReturnDefaultValue("USER_SERVICE_PORT", 8109))
	cfg.ContentServiceHost = cast.ToString(getOrReturnDefaultValue("CONTENT_SERVICE_HOST", "localhost"))
	cfg.ContentServicePort = cast.ToInt(getOrReturnDefaultValue("CONTENT_SERVICE_PORT", 7004))
	cfg.AnalyticServiceUrl = cast.ToString(getOrReturnDefaultValue("ANALYTIC_SERVICE_URL", "127.0.0.1:8081"))

	// redis
	cfg.RedisHost = cast.ToString(getOrReturnDefaultValue("REDIS_HOST", "localhost"))
	cfg.RedisPassword = cast.ToString(getOrReturnDefaultValue("REDIS_PASSWORD", "redis"))
	cfg.RedisPort = cast.ToInt(getOrReturnDefaultValue("REDIS_PORT", 6379))
	cfg.RedisTimeoutInSeconds = cast.ToInt(getOrReturnDefaultValue("REDIS_TIMEOUT_IN_SECONDS", 120))

	// minio
	cfg.BucketName = cast.ToString(getOrReturnDefaultValue("MINIO_BUCKET_NAME", "ekadastr"))
	cfg.FilesBucketName = cast.ToString(getOrReturnDefaultValue("MINIO_FILE_BUCKET_NAME", "files"))
	cfg.JackBucketName = cast.ToString(getOrReturnDefaultValue("MINIO_ZONES_BUCKET_NAME", "enchanted"))
	cfg.MinioDomain = cast.ToString(getOrReturnDefaultValue("MINIO_DOMAIN", "test-cdn.yerelektron.uz"))
	cfg.MinioAccessKeyID = cast.ToString(getOrReturnDefaultValue("MINIO_ACCESS_KEY", "8DbGdJfNjQmSqVsXv2x4z7C9EbHeKgNkRnTrWtYv3y5A7DaFcJfMhPmSpUrXuZw3z6B8EbGdJgNjQmTqVsXv2x4A7C"))
	cfg.MinioSecretAccesKey = cast.ToString(getOrReturnDefaultValue("MINIO_SECRET_KEY", "QmSpUsXuZx4z6B9EbGdKgNjQnTqVtYv2x5A7C9FcHeKhPkRpUrWtZw3y5B8DaFdJfMjQmSpVsXuZx4z6B9EbGeKgNj"))
	cfg.MinioLocation = cast.ToString(getOrReturnDefaultValue("MINIO_LOCATION", "us-east-1"))

	// JWT
	cfg.AccessTokenExpireDuration = cast.ToUint(getOrReturnDefaultValue("ACCESS_TOKEN_DURATION_HOURS", 1))
	cfg.RefreshTokenExpireDuration = cast.ToUint(getOrReturnDefaultValue("REFRESH_TOKEN_DURATION_HOURS", 1))

	// keys
	cfg.LoginSecretAccessKey = cast.ToString(getOrReturnDefaultValue("LOGIN_ACCESS_SECRET_KEY", "dWRldnMgZGV2ZWxvcGVkIGVsZWN0cm9uIGthZGFzdHIK"))
	cfg.LoginSecretRefreshKey = cast.ToString(getOrReturnDefaultValue("LOGIN_REFRESH_SECRET_KEY", "ZWxlY3Ryb24ga2FkYXN0ciAgLT4gdWRldnMgZGV2ZWxvcGVkCg"))
	return cfg
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	source := ".env"
	envErr := godotenv.Load(source)
	if envErr != nil {
		fmt.Println("Could not find " + source + " file")
	}

	if val := os.Getenv(key); val != "" {
		return val
	}

	return defaultValue
}
