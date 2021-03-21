package utils

import (
	"bytes"
	"encoding/json"
)

const (
	empty = ""
	tab   = "	"
)

func FormatJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}

	return buffer.String(), nil
}
