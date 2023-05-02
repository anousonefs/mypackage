package myrandom

import (
	"math/rand"
	"time"
)

var MIN = 0
var MAX = 100
var SEED = time.Now().Unix()

func Random(total int) int {
	result := 0
	rand.Seed(SEED)
	for i := 0; i < total; i++ {
		myrand := random(MIN, MAX)
		result += myrand
	}
	return 0
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func Hello() {
	println("hello anousone")
}
