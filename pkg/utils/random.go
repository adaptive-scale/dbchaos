package utils

import (
	"math/rand"
	"time"
)

type Charset string

const (
	CharsetAlphaNumeric Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharsetAlphabet     Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharsetLower        Charset = "abcdefghijklmnopqrstuvwxyz"
	CharsetUpper        Charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func RandomString(charset Charset, length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
