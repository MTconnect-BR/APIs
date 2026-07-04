package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateID() string {
	return uuid.New().String()
}

func GeneratePrefixedID(prefix string) string {
	return fmt.Sprintf("%s_%s", prefix, uuid.New().String()[:8])
}
