package utils

import (
	"encoding/json"
	"io/ioutil"
)

func DecodeJsonFromFile(i interface{}, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, i)
}
