package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/dearing/numbers/english"
)

func main() {
	rand.Seed(time.Now().Unix())
	test := rand.Intn(math.MaxInt) + -rand.Intn(math.MaxInt)
	println(test, "is", english.NumberToEnglish(test))
}
