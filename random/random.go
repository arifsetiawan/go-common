package random

import "math/rand"

// Credits to icza of https://stackoverflow.com/a/31832326/1161743 for the original code.
// IMPORTANT. Seed rand first to ensure you don't get the same string on every code run (initialize):
// rand.Seed(time.Now().UTC().UnixNano())
// rand.Seed(time.Now().Unix())

// Character sets for random string generation
const (
	CharSetAlphabet           = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharSetAlphaNumeric       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	CharSetAlphaNumericSymbol = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&()*+,-./:;<=>?@[]^_`{|}~"
)

var defaultLetterBytes = CharSetAlphabet

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Generator is the instance that generates random strings from the given charset
type Generator struct {
	charset string
}

// NewGenerator returns a new Generator instance
func NewGenerator(charset string) *Generator {
	return &Generator{
		charset: charset,
	}
}

// GenerateRandomString generates a length-n random string from the given character set
func (g *Generator) GenerateRandomString(n int) string {
	b := make([]byte, n)

	// A rand.Int63() generates 63 random bits, enough for letterIdxMax letters!
	for i, cache, remain := n-1, rand.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = rand.Int63(), letterIdxMax
		}

		if idx := int(cache & letterIdxMask); idx < len(g.charset) {
			b[i] = g.charset[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// GenerateAlphabet is
func GenerateAlphabet(n int) string {
	return NewGenerator(CharSetAlphabet).GenerateRandomString(n)
}

// GenerateAlphaNumeric is
func GenerateAlphaNumeric(n int) string {
	return NewGenerator(CharSetAlphaNumeric).GenerateRandomString(n)
}

// GenerateAlphaNumericSymbol is
func GenerateAlphaNumericSymbol(n int) string {
	return NewGenerator(CharSetAlphaNumericSymbol).GenerateRandomString(n)
}
