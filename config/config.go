package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload" // load .env file automatically
	"github.com/spf13/cast"
)

var (
	instance *Configuration
	once     sync.Once
)

// Config ...
func Config() *Configuration {
	once.Do(func() {
		instance = load()
	})
	return instance
}

// Configuration ...
type Configuration struct {
	Server      Server
	JWT         JWT
	Postgres    Database
	Minio       MinioStore
	APIProperty APISettings
	Security    Security
	Credential Credential
}
type Server struct {
	AppName     string
	AppVersion  string
	Environment string
	ServerPort  int
	AppURL      string
	ServerHost  string
	MigrationUp bool
	StoreUrl    string
}
type JWT struct {
	BasicAuthUsername         string
	BasicAuthPassword         string
	JWTSecretKey              string
	JWTSecretKeyExpireMinutes int
	JWTRefreshKey             string
	JWTRefreshKeyExpireHours  int
}
type Database struct {
	Host     string
	Port     int
	Name     string
	UserName string
	Password string
	SSLMode  string
}
type MinioStore struct {
	MinioAccessKeyID      string
	MinioSecretKey        string
	MinioEndpoint         string
	MinioBucketName       string
	MinioPublicBucketName string
	MinioLocation         string
	MinioUseSSL           bool
}
type APISettings struct {
	DefaultOffset   string
	DefaultLimit    string
	DefaultPage     string
	DefaultPageSize string
}
type Security struct {
	HashKey string
}

type Credential struct {
	ConsumerKey string
	ConsumerSecret string
}

func load() *Configuration {
	return &Configuration{
		Server: Server{
			AppURL:      cast.ToString(getOrReturnDefault("APP_URL", "localhost:9000")),
			AppName:     cast.ToString(getOrReturnDefault("APP_NAME", "ProPay")),
			AppVersion:  cast.ToString(getOrReturnDefault("APP_VERSION", "1.0")),
			ServerHost:  cast.ToString(getOrReturnDefault("SERVER_HOST", "localhost")),
			ServerPort:  cast.ToInt(getOrReturnDefault("SERVER_PORT", "9000")),
			Environment: cast.ToString(getOrReturnDefault("ENVIRONMENT", "production")),
			MigrationUp: cast.ToBool(getOrReturnDefault("MIGRATION_UP", false)),
			StoreUrl:    cast.ToString(getOrReturnDefault("STORE_URL", "")),
		},
		JWT: JWT{
			JWTSecretKey:              cast.ToString(getOrReturnDefault("JWT_SECRET_KEY", "")),
			JWTSecretKeyExpireMinutes: cast.ToInt(getOrReturnDefault("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", 720)),
			JWTRefreshKey:             cast.ToString(getOrReturnDefault("JWT_REFRESH_KEY", "")),
			JWTRefreshKeyExpireHours:  cast.ToInt(getOrReturnDefault("JWT_REFRESH_KEY_EXPIRE_HOURS_COUNT", 24)),
		},
		Postgres: Database{
			Host:     cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost")),
			Port:     cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432)),
			Name:     cast.ToString(getOrReturnDefault("POSTGRES_DB", "")),
			UserName: cast.ToString(getOrReturnDefault("POSTGRES_USER", "")),
			Password: cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "")),
			SSLMode:  cast.ToString(getOrReturnDefault("POSTGRES_SSLMODE", "disable")),
		},
		Minio: MinioStore{
			MinioAccessKeyID:      cast.ToString(getOrReturnDefault("MINIO_ACCESS_KEY_ID", "")),
			MinioSecretKey:        cast.ToString(getOrReturnDefault("MINIO_SECRET_KEY", "")),
			MinioEndpoint:         cast.ToString(getOrReturnDefault("MINIO_ENDPOINT", "")),
			MinioBucketName:       cast.ToString(getOrReturnDefault("MINIO_BUCKET_NAME", "")),
			MinioPublicBucketName: cast.ToString(getOrReturnDefault("MINIO_PUBLIC_BUCKET_NAME", "")),
			MinioLocation:         cast.ToString(getOrReturnDefault("MINIO_LOCATION", "")),
			MinioUseSSL:           cast.ToBool(getOrReturnDefault("MINIO_USE_SSL", false)),
		},
		APIProperty: APISettings{
			DefaultOffset:   cast.ToString(getOrReturnDefault("DEFAULT_OFFSET", 0)),
			DefaultLimit:    cast.ToString(getOrReturnDefault("DEFAULT_LIMIT", 10)),
			DefaultPage:     cast.ToString(getOrReturnDefault("DEFAULT_PAGE", 1)),
			DefaultPageSize: cast.ToString(getOrReturnDefault("DEFAULT_PAGE_SIZE", 10)),
		},
		Security: Security{
			HashKey: cast.ToString(getOrReturnDefault("HASH_KEY", "")),
		},
		Credential: Credential{
			ConsumerKey: cast.ToString(getOrReturnDefault("CONSUMERKEY", "")),
			ConsumerSecret: cast.ToString(getOrReturnDefault("CONSUMERSECRET", "")),
		},
	}
}
func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}
	return defaultValue
}
