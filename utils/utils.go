package utils

import (
	"os"
)


// GetEnv get environment vars with deafult value
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
			value = fallback
	}
	return value
}

// WriteFile is write something into file
func WriteFile(name string, content []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	return err
}