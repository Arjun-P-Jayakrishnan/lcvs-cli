package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Dir
func CreateDir(path string) error {
	return os.Mkdir(path, 0755)
}

// Path
func PathJoin(parts ...string) string {
	return filepath.Join(parts...)
}

func AbsPath(path string) (string, error) {
	return filepath.Abs(path)
}

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

func ReadJSON(path string, out any) error {

	f, err := os.Open(path)
	if err != nil {
		return err
	}

	defer f.Close()

	return json.NewDecoder(f).Decode(out)
}

// Plain Text
func ReadFile(path string) ([]byte, error) {

	return ioutil.ReadFile(path)
}

func WriteFile(path string, data []byte) error {
	
	return ioutil.WriteFile(path, data, 0755)
}
