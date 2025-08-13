package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func TokenURLSafe64() (string, error) {
	b := make([]byte, 48)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	token := base64.RawURLEncoding.EncodeToString(b)
	return token, nil
}

func main() {
	token, err := TokenURLSafe64()
	if err != nil {
		panic(err)
	}
	fmt.Println(token)
}

