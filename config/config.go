package config

import (
	"kage/utils/logger"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var (
	Config *ServerConfig
	log    = logger.NewLogger().WithPrefix("Config")
)

func init() {
	godotenv.Load()

	Config = &ServerConfig{
		Port:  getEnvAsInt("PORT", 3000),
		Debug: getEnvAsBool("DEBUG", false),
	}

	if Config.Debug {
		log.WithTimestamp().WithTimeFormat("02/01/2006 03:04:05 PM")
	}

	if Config.Port == 0 {
		log.Fatalf("Port is not set in the environment variables")
	}

	log.Successf("Configuration loaded successfully")
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}

	return strings.TrimSpace(value)
}

func getEnvAsInt(key string, defaultValue int) int {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}

	return boolValue
}
