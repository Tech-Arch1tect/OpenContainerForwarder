package config

import (
	"os"
	"strconv"
)

// getDefaultStr returns the value of a string environment variable or a default value if the variable is not set
func getDefaultStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// getDefaultInt returns the value of a integer environment variable or a default value if the variable is not set
func getDefaultInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	v, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return v
}

// getDefaultBool returns the value of a boolean environment variable or a default value if the variable is not set
func getDefaultBool(key string, defaultValue bool) bool {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	v, err := strconv.ParseBool(value)
	if err != nil {
		return defaultValue
	}
	return v
}
