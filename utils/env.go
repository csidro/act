package utils

import (
	"os"
	"strconv"
)

func GetEnvAsNumber(key string, defaultValue int) int {
	value, present := os.LookupEnv(key)
	if !present {
		return defaultValue
	}

	asNumber, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return asNumber
}

func GetEnv(key string, defaultValue string) string {
	value, present := os.LookupEnv(key)
	if !present {
		return defaultValue
	}

	return value
}
