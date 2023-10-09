package util

import (
	json "github.com/json-iterator/go"
)

func Transform(in interface{}, dst interface{}) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}
