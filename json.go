package main

import "encoding/json"

func SimpleMarshal(val interface{}) ([]byte, error) {
	return json.Marshal(val)
}

func SimpleUnmarshal(data []byte, val interface{}) error {
	return json.Unmarshal(data, val)
}