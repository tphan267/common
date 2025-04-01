package utils

import "encoding/json"

func Copy(src any, dest any) error {
	if data, err := json.Marshal(src); err != nil {
		return err
	} else {
		return json.Unmarshal(data, dest)
	}
}
