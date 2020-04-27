package main

import "github.com/intel-go/fastjson"

func FastMarshal(val interface{}) ([]byte, error) {
	return fastjson.Marshal(val)
}

func FastUnmarshal(data []byte, val interface{}) error {
	return fastjson.Unmarshal(data, val)
}