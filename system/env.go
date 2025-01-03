package system

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var loadEnv = false

// Config function to get value from env file
// perhaps not the best implementation
func Env(key string, defaultValue ...string) string {
	// load .env file
	if loadEnv == false {
		if _, err := os.Stat(".env"); err != nil {
			fmt.Printf("Config .env file does not exist! %s", err.Error())
		} else {
			if err := godotenv.Load(".env"); err != nil {
				fmt.Printf("Error loading .env file! %s", err.Error())
			}
		}
		loadEnv = true
	}

	val := os.Getenv(key)
	if val == "" && len(defaultValue) > 0 {
		return defaultValue[0]
	}

	return val
}

func EnvInt(key string, defaultValue ...int) int {
	if val := Env(key); val != "" {
		intVal, err := strconv.Atoi(val)
		if err == nil {
			return intVal
		}
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return 0
}
