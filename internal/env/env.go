package env

import (
	"log"
	"os"
	"strconv"
	"time"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

func GetDuration(key, fallback string) time.Duration {
	fbVal, err := time.ParseDuration(fallback)
	if err != nil {
		log.Fatalf("Invalid fallback provided as: %s", fallback)
	}

	val, ok := os.LookupEnv(key)
	if !ok {
		return fbVal
	}

	duration, err := time.ParseDuration(val)
	if err != nil {
		return fbVal
	}

	return duration
}
