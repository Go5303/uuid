package main

import (
	"math/rand"
	"time"
)

const StaticStr = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

var str string

func Uuid() string {
	strLen := len(StaticStr)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i <= 100; i++ {
		number := r.Intn(strLen)
		str += StaticStr[number : number+1]
	}
	return str
}
