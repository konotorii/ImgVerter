package util

import (
	"os"
	"strconv"
	"strings"
)

type ConfigStruct struct {
	CookieSecret   string
	Port           int
	RedisHost      string
	RedisPass      string
	RedisDb        int
	Domain         string
	UploadKey      string
	UploadSettings UploadSettings
	PublicFolder   string
}

type UploadSettings struct {
	AllowedFileTypes []string
	MaxFileSize      int // in MB !!!
}

func ConfigInit() {
	Config = &ConfigStruct{
		CookieSecret: getEnv("SECRET", ""),
		Port:         getEnvAsInt("PORT", 3000),
		RedisHost:    getEnv("REDIS_HOST", "localhost:6379"),
		RedisPass:    getEnv("REDIS_PASS", ""),
		RedisDb:      getEnvAsInt("REDIS_DB", 1),
		Domain:       getEnv("DOMAIN", "localhost"),
		UploadKey:    getEnv("UPLOAD_KEY", ""),
		PublicFolder: getEnv("PUBLIC_FOLDER", "/public/"),
		UploadSettings: UploadSettings{
			AllowedFileTypes: strings.Split(getEnv("UPLOAD_ALLOWED_FILE_TYPES", "png,jpg,jpeg,gif,mp4"), ","),
			MaxFileSize:      getEnvAsInt("UPLOAD_MAX_FILE_SIZE", 5),
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
