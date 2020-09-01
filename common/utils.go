package common

import (
	"bytes"
	"log"
	"encoding/json"
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

// ConverMapToJSONPayload returns  bytes buffer 
func ConverMapToJSONPayload(payload map[string]interface{}) (*bytes.Buffer, error) {
	bytesRepresentationPayload, err := json.Marshal(payload)
	if err != nil {
		return  nil, err
	}
	buff :=bytes.NewBuffer(bytesRepresentationPayload)
	return buff, nil
}
