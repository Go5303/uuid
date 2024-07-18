package uuid

import (
	"math/rand"
	"time"
)

const StaticStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func Uuid() string {
	var str string
	strLen := len(StaticStr)
	for i := 0; i <= 100; i++ {
		number := r.Intn(strLen)
		str += StaticStr[number : number+1]
	}
	return str
}
