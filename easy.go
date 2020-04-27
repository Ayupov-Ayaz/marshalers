package main

import "github.com/mailru/easyjson"

func EasyMarshal(m easyjson.Marshaler) ([]byte, error) {
	return easyjson.Marshal(m)
}

func EasyUnmarshal(data []byte, m easyjson.Unmarshaler) error {
	return easyjson.Unmarshal(data, m)
}