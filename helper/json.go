package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, data interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	err := encoder.Encode(&response)
	PanicIfError(err)
}
