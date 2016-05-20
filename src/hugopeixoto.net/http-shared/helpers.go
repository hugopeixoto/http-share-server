package main

import "encoding/json"
import "net/http"

type JSONError struct {
	Detail string `json:"detail"`
}

type ErrorPayload struct {
	Errors []JSONError `json:"errors"`
}

type DataPayload struct {
	Data interface{} `json:"data"`
}

func Error(w http.ResponseWriter, code int, err error) {
	json, _ := json.Marshal(
		ErrorPayload{[]JSONError{JSONError{err.Error()}}})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(json)
}

func Data(w http.ResponseWriter, code int, obj interface{}) {
	json, _ := json.Marshal(DataPayload{obj})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(json)
}
