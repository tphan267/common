package utils

import "encoding/json"

func Copy(src interface{}, dest interface{}) error {
	if data, err := json.Marshal(src); err != nil {
		return err
	} else {
		return json.Unmarshal(data, dest)
	}
}
