package common

import (
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"encoding/json"
)

// NewHTTPError example
func NewHTTPError(ctx *gin.Context, status int, err error) {
	er := HTTPError{
		Code:    status,
		Message: err.Error(),
	}
	ctx.JSON(status, er)
}

// HTTPError example
type HTTPError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"Request Failure"`
}

// MakeHTTPRequestWithJSONPayload makes json payload
func MakeHTTPRequestWithJSONPayload(payload map[string]interface{}, URI string) {
	payloadBuffer, err := ConverMapToJSONPayload(payload)
	if(err != nil){
		log.Fatalln(err);
	}
	resp, err := http.Post(URI, "application/json; charset=UTF-8", payloadBuffer)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	log.Println(result)
	log.Println("data: ", result["data"])
}
