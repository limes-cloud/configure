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

// Diff 计算两个数组中的key变更元素
func Diff(cur, old []string) []string {
	set := make(map[string]bool)
	n := make([]string, 0)
	for _, v := range old {
		set[v] = true
	}

	for _, v := range cur {
		if !set[v] {
			n = append(n, v)
		}
	}
	return n
}

// Intersect 计算两个数组中交叉的数组
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]bool)
	n := make([]string, 0)
	for _, v := range slice1 {
		m[v] = true
	}

	for _, v := range slice2 {
		if m[v] {
			n = append(n, v)
		}
	}
	return n
}
