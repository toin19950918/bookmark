package md5

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// New a md5 hash
func New(data interface{}, secret string) string {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	s := string(jsonStr) + secret
	result := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", result)
}

// New a md5 hash
func Concat(data string, secret string) string {
	s := data + secret
	result := md5.Sum([]byte(s))
	//KeyString := hex.EncodeToString(result[:])
	//fmt.Println("key is ", KeyString)
	return fmt.Sprintf("%x", result)
}
