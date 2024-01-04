package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	json "github.com/json-iterator/go"
)

func Transform(in any, dst any) error {
	b, err := json.Marshal(in)
	if err != nil {
		return err
	}
	return json.Unmarshal(b, dst)
}

func MD5(in []byte) string {
	sum := md5.Sum(in)
	return hex.EncodeToString(sum[:])
}

func MD5ToUpper(in []byte) string {
	return strings.ToUpper(MD5(in))
}

func MarshalString(in any) string {
	b, _ := json.Marshal(in)
	return string(b)
}
