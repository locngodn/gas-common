package util

import (
	mRand "math/rand"
	"net/http"
	"strings"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

var numbers = []rune("0123456789")

func GetRandomString(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = letters[mRand.Intn(len(letters))]
	}
	return string(b)
}

func ToChar(i int) rune {
	return rune('A' - 1 + i)
}

// getRandNum returns a random number
func GetRandNumber(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = numbers[mRand.Intn(len(numbers))]
	}
	return string(b)
	//nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	//if e != nil {
	//	return "", e
	//}
	//return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func GetHttpAccessToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer")
	return strings.TrimSpace(splitToken[1])
}
