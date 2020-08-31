package common

import (
	"log"

	"github.com/google/uuid"
)

// GetNewUUID retuns uuid
func GetNewUUID() uuid.UUID {
	return uuid.New()
}

// ConvertStringToID returns Id
func ConvertStringToID(s string) uuid.UUID {
	uuid, err := uuid.Parse(s)
	if err != nil {
		log.Fatal(err.Error())
	}
	return uuid
}
