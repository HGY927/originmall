package utils

import "github.com/google/uuid"

func UuidValue() string {
	return uuid.New().String()

}
