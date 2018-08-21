package utils

import "math/rand"

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// AlphaNumString generate an alphanumeric string with n length
func AlphaNumString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes))]
	}
	return string(b)
}

// LetterString generates a letter only strings with n length
func LetterString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
