package util

import "github.com/google/uuid"

type UUIDString string

func ValidUUID(id string) bool {
	if _, err := uuid.Parse(id); err != nil {
		return false
	}
	return true
}
