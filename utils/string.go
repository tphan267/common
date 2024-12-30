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

func GenerateID() (string, error) {
	id := make([]byte, idLength)
	// Generate prefix
	for i := 0; i < idLength; i++ {
		char, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabets))))
		if err != nil {
			return "", err
		}
		id[i] = alphabets[char.Int64()]
	}
	return string(id), nil
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
