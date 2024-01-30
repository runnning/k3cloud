package object

import (
	"encoding/json"
	"io/fs"
	"os"
)

func JsonEncode(v any) (string, error) {
	buffer, err := json.Marshal(v)

	if err != nil {
		return "", err
	}
	return string(buffer), nil
}

func JsonDecode(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func JsonEscape(str string) (string, error) {
	b, err := json.Marshal(str)
	if err != nil {
		return "", err
	}
	return string(b[1 : len(b)-1]), err
}

func SaveObjectToFile(obj any, filePath string, perm fs.FileMode) (err error) {

	buff, err := json.MarshalIndent(obj, "", " ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, buff, perm)

	return err
}
