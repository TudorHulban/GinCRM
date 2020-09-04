package authentication

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const asciiPartLen = 7

// GenerateRandomString To be used for password salt.
func generateRandomString(targetLength int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]string, targetLength)

	randInt := func(min, max int) int {
		return min + rand.Intn(max-min)
	}
	for k := range result {
		result[k] = string(byte(randInt(65, 90)))
	}
	return strings.Join(result, "")
}

func generateSessionID() string {
	return strconv.FormatInt(time.Now().UnixNano(), 10) + generateRandomString(asciiPartLen)
}
