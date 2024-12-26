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
func Env(key string, _default ...string) string {
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
	if val == "" && len(_default) > 0 {
		return _default[0]
	}

	return val
}

func EnvAsInt(key string, _default int) int {
	if val := Env(key); val != "" {
		iVal, err := strconv.Atoi(val)
		if err != nil {
			return _default
		}
		return iVal
	}

	return _default
}
