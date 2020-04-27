package main

import (
	"github.com/pquerna/ffjson/ffjson"
)

func FfJsonMarshal(val interface{}) ([]byte, error) {
	return ffjson.Marshal(val)
}

func FfJsonUnmarshal(data []byte, val interface{}) error {
	return ffjson.Unmarshal(data, val)
}