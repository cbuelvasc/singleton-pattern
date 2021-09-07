package main

import (
	"fmt"

	"github.com/cbuelvasc/singleton-pattern/counter"
)

func main() {
	counterOne := counter.GetCounter()
	currentCount := counterOne.AddOne()
	fmt.Println(currentCount)

	counterTwo := counter.GetCounter()
	currentCount = counterTwo.AddOne()
	fmt.Println(currentCount)

	counterThree := counter.GetCounter()
	currentCount = counterThree.SubtractOne()
	fmt.Println(currentCount)
}
