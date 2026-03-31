package uuid

import (
	"math/rand"
	"strings"
)

const StaticStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func Uuid() string {
	var b strings.Builder
	b.Grow(100)
	strLen := len(StaticStr)
	for i := 0; i < 100; i++ {
		b.WriteByte(StaticStr[rand.Intn(strLen)])
	}
	return b.String()
}
