package authentication

import (
	"math/rand"

	"github.com/TudorHulban/GinCRM/cmd/setup"
)

/*
File provides global helpers for salt and hashing.
*/

// GenerateSALT Generates salt based on settings value.
func GenerateSALT() string {
	return randomString(setup.LenSALT)
}

func randomString(length int) string {
	letterBytes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	result := make([]byte, length)
	for i := range result {
		result[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(result)
}
