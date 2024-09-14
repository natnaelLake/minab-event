package utilService

import (
	"math/rand"
	"time"
)

func PublicID()  string {
	rand.Seed(time.Now().UnixNano())
	length := 18
	chars := "abcdefghijklmnopqrstuvwxyz_0123456789-ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	word := ""
	for i := 0; i < length; i++ {
		index := rand.Intn(len(chars))
		word += string(chars[index])
	}
	return word
}