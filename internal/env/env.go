package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	val, done := os.LookupEnv(key)
	if !done {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, done := os.LookupEnv(key)
	if !done {
		return fallback
	}

	v, err := strconv.Atoi(val)

	if err != nil {
		return fallback
	}

	return v
}
