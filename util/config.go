package util

import (
	"os"
	"strconv"
	"strings"
)

type ConfigStruct struct {
	TimeZone         string
	CookieSecret     string
	Port             int
	UseRedis         bool
	DatabaseDriver   string
	RedisHost        string
	RedisPass        string
	RedisDb          int
	Domain           string
	UploadKey        string
	UploadSettings   UploadSettings
	DatabaseSettings DatabaseSettings
	PublicFolder     string
}

type UploadSettings struct {
	AllowedFileTypes     []string
	MaxFileSize          int // in MB !!!
	EnableWebpConversion bool
}

type DatabaseSettings struct {
	Path string /// relative to application
	Host string
	Port int
	User string
	Pass string
	Name string
}

func ConfigInit() {
	Config = &ConfigStruct{
		TimeZone:       getEnv("TZ", "UTC"),
		CookieSecret:   getEnv("SECRET", ""),
		UseRedis:       getEnvAsBool("USE_REDIS", false),
		DatabaseDriver: getEnv("DATABASE_DRIVER", "SQLITE"),
		Port:           getEnvAsInt("PORT", 3000),
		RedisHost:      getEnv("REDIS_HOST", "localhost:6379"),
		RedisPass:      getEnv("REDIS_PASS", ""),
		RedisDb:        getEnvAsInt("REDIS_DB", 1),
		Domain:         getEnv("DOMAIN", "localhost"),
		UploadKey:      getEnv("UPLOAD_KEY", ""),
		PublicFolder:   getEnv("PUBLIC_FOLDER", "/public/"),
		UploadSettings: UploadSettings{
			AllowedFileTypes:     strings.Split(getEnv("UPLOAD_ALLOWED_FILE_TYPES", "png,jpg,jpeg,gif,mp4"), ","),
			MaxFileSize:          getEnvAsInt("UPLOAD_MAX_FILE_SIZE", 5),
			EnableWebpConversion: getEnvAsBool("UPLOAD_WEBP_CONVERSION", true),
		},
		DatabaseSettings: DatabaseSettings{
			Path: "./private/" + getEnv("DATABASE_PATH", "imgverter.db"),
			Host: getEnv("DATABASE_HOST", "localhost"),
			Port: getEnvAsInt("DATABASE_PORT", 0),
			User: getEnv("DATABASE_USER", "admin"),
			Pass: getEnv("DATABASE_PASS", "admin"),
			Name: getEnv("DATABASE_NAME", "imgverter"),
		},
	}
}

var Config *ConfigStruct

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseBool(valueStr); err == nil {
		return value
	}

	return defaultVal
}
