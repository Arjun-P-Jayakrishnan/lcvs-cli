package utils

import (
	"encoding/json"
	"os"
)

// JSON
func WriteJSON(path string, data any) error {
	f, err := os.Create(path)

	if err != nil {
		return err
	}

	defer f.Close()

	enc := json.NewEncoder(f)
	enc.SetIndent("", " ")
	return enc.Encode(data)
}

func ReadJSON(path string, out interface{}) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(out)
}
