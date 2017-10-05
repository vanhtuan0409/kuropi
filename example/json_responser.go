package main

import (
	"encoding/json"
	"net/http"
)

type JsonResponser struct{}

func (rs *JsonResponser) Handle(rw http.ResponseWriter, result interface{}, err error) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	var data map[string]interface{}
	if err != nil {
		data = map[string]interface{}{
			"ok":    false,
			"error": err.Error(),
		}
	} else {
		data = map[string]interface{}{
			"ok":   true,
			"data": result,
		}
	}
	jData, _ := json.Marshal(data)
	rw.Write(jData)
}
