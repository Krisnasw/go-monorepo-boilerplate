package helpers

import (
	"math/rand"
	"time"
)

const randomStringLetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	randomStringletterIdxBits = 6
	randomStringletterIdxMask = 1<<randomStringletterIdxBits - 1
	randomStringletterIdxMax  = 63 / randomStringletterIdxBits
)

func RandomString(n int) string {
	// from : https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go
	b := make([]byte, n)
	for i, cache, remain := n-1, rand.Int63(), randomStringletterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), randomStringletterIdxMax
		}
		if idx := int(cache & randomStringletterIdxMask); idx < len(randomStringLetterBytes) {
			b[i] = randomStringLetterBytes[idx]
			i--
		}
		cache >>= randomStringletterIdxBits
		remain--
	}

	return string(b)
}

const letterNumberBytes = "0123456789"

func RandomStringIntOnly(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterNumberBytes[rand.Int63()%int64(len(letterNumberBytes))]
	}
	return string(b)
}
