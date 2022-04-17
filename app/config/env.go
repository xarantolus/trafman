package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	GitHubToken             string
	AppPort                 string
	DisableBackgroundChecks bool
	Debug                   bool
	DB                      DB
}

type DB struct {
	Host     string
	Port     string
	User     string
	Password string

	DBName string
}

func getOrDefault(key, def string) string {
	val := os.Getenv(key)
	if strings.TrimSpace(val) == "" {
		return def
	}
	return val
}

func getBoolean(key string, def bool) (bool, error) {
	env := os.Getenv(key)
	if strings.TrimSpace(env) == "" {
		return def, nil
	}

	envVal, err := strconv.ParseBool(env)
	if err != nil {
		return false, fmt.Errorf("parsing %s environment variable: %s", key, err.Error())
	}

	return envVal, nil
}

func FromEnvironment() (cfg Config, err error) {
	var ghToken = os.Getenv("GITHUB_TOKEN")
	if strings.TrimSpace(ghToken) == "" {
		err = fmt.Errorf("no GITHUB_TOKEN env variable available")
		return
	}

	disableBGChecks, err := getBoolean("APP_DISABLE_BACKGROUND_CHECKS", false)
	if err != nil {
		return
	}
	enableDebugMode, err := getBoolean("APP_DEBUG", false)
	if err != nil {
		return
	}

	return Config{
		DB: DB{
			Host:     getOrDefault("DB_HOST", "postgres"),
			Port:     getOrDefault("DB_PORT", "5432"),
			User:     getOrDefault("DB_USER", "postgres"),
			Password: getOrDefault("DB_PASSWORD", "postgres"),
			DBName:   getOrDefault("DB_NAME", "trafmon_database"),
		},

		Debug:                   enableDebugMode,
		DisableBackgroundChecks: disableBGChecks,
		AppPort:                 getOrDefault("APP_PORT", "9319"),
		GitHubToken:             ghToken,
	}, nil
}
