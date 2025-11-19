package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	Port           uint32
	RootPath       string
	MaxUploadBytes uint64
	UsageCacheTTL  uint32
	LogLevel       string
	GinMode        string
}

func New() *Config {
	return &Config{}
}

func (config *Config) Load() {
	portStr := os.Getenv("PORT")
	rootPath := os.Getenv("ROOT_PATH")
	maxUploadBytesStr := os.Getenv("MAX_UPLOAD_BYTES")
	usageCacheTTLStr := os.Getenv("USAGE_CACHE_TTL")
	logLevel := os.Getenv("LOG_LEVEL")
	ginMode := os.Getenv("GIN_MODE")

	port, err := strconv.Atoi(portStr)

	if err != nil {
		log.Panic("PORT must be an integer")
		return
	}

	config.Port = uint32(port)
	config.RootPath = rootPath

	maxUploadBytes, err := strconv.Atoi(maxUploadBytesStr)

	if err != nil {
		log.Panic("MAX_UPLOAD_BYTES must be an integer")
		return
	}

	config.MaxUploadBytes = uint64(maxUploadBytes)

	usageCacheTTL, err := strconv.Atoi(usageCacheTTLStr)

	if err != nil {
		log.Fatal("USAGE_CACHE_TTL must be an integer")
		return
	}

	config.UsageCacheTTL = uint32(usageCacheTTL)
	config.GinMode = ginMode
	config.LogLevel = logLevel

}
