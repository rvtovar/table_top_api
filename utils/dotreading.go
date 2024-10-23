package utils

import (
	"errors"
	"github.com/joho/godotenv"
	"os"
)

func OpenSalt() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}
	salt := os.Getenv("salt")
	if salt == "" {
		return "", errors.New("Salt not found")
	}

	return salt, nil
}

func OpenDBStr() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", errors.New("Error loading .env file")
	}
	dbStr := os.Getenv("db")
	return dbStr, nil
}
