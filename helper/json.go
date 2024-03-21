package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result interface{}) error {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		return err
	}
	return nil
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	encoder.Encode(response)
}

func WriteErrToResponseBody(writer http.ResponseWriter, err error) {
	writer.Header().Add("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	encoder := json.NewEncoder(writer)
	encoder.Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
