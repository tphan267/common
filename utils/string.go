package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"
)

const (
	idLength  = 8
	alphabets = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateRandomString(length int, seeds ...string) (string, error) {
	id := make([]byte, length)
	seed := alphabets
	if len(seeds) > 0 {
		seed = seeds[0]
	}
	// Generate prefix
	for i := 0; i < length; i++ {
		char, err := rand.Int(rand.Reader, big.NewInt(int64(len(seed))))
		if err != nil {
			return "", err
		}
		id[i] = seed[char.Int64()]
	}

	return string(id), nil
}

func GenerateID() (string, error) {
	return GenerateRandomString(idLength)
}

func GeneratePassword(length int) (string, error) {
	return GenerateRandomString(length, alphabets+"!@#$%^&*()_+")
}

func StringToInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}

func ToString(data interface{}) string {
	switch data.(type) {
	case int:
		return fmt.Sprintf("%d", data)
	case float32:
	case float64:
		return fmt.Sprintf("%f", data)
	case string:
		return data.(string)
	default:
	}
	text, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return ""
	}
	return string(text)
}

func HashKey(key string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(key)))
}
