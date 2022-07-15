package main

import (
	"encoding/base64"
	"fmt"
)

func reverse(str string) (result string) {
	for _, char := range str {
		result = string(char) + result
	}
	return
}

func main() {
	var whatIsIt string

	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"

	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = reverse(string(sd))
	fmt.Println(whatIsIt)
