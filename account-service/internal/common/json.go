package common

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

func BindJSON(source io.Reader, data interface{}) error {
	decoder := json.NewDecoder(source)
	return decoder.Decode(data)
}

func ToJSON(res http.ResponseWriter, status int, data interface{}) {
	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)

	if err := encoder.Encode(data); err != nil {
		res.Header().Set("Content-Type", "text/plain")
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
	}

	res.Header().Set("Content-Type", "application/json;charset=UTF-8")
	res.Header().Set("Server", "Account/1.0")
	res.WriteHeader(status)
	res.Write(buffer.Bytes())
}
