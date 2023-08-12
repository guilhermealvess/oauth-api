package entity

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func removeSpecialCharacters(input string) string {
	reg := regexp.MustCompile("[^a-zA-Z0-9 ]+")
	return reg.ReplaceAllString(input, "")
}

func slugFy(s string) string {
	slug := strings.ToLower(s)
	slug = strings.ReplaceAll(slug, " ", "-")
	slug = strings.ReplaceAll(slug, "_", "-")
	return slug
}

func randomString(l int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, l)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func calculateSHA256(secret, salt string, round int) string {
	for i := 0; i < round; i++ {
		secret = secret + salt

		hash := sha256.New()
		hash.Write([]byte(secret))
		hashSum := hash.Sum(nil)
		secret = hex.EncodeToString(hashSum)
	}
	return secret

}
