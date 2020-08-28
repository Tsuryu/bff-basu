package utils

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/Tsuryu/bff-basu/models"
)

// GetResponseBody : handle response from server
func GetResponseBody(body []byte, statusCode int, entity interface{}) error {
	if statusCode != 200 {
		model, _ := getErrorResponse(body)
		return errors.New(model.Result)
	}

	err := json.Unmarshal(body, &entity)
	if err != nil {
		log.Println("Failed to unmarshal")
		return err
	}

	return nil
}

func getErrorResponse(body []byte) (model models.ResponseError, err error) {
	err = json.Unmarshal(body, &model)
	if err != nil {
		log.Println("Failed to unmarshal")
		return model, err
	}

	return model, nil
}
