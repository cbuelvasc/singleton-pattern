package main

import (
	"fmt"

	singleton "github.com/cbuelvasc/singleton-pattern/counter"
)

func main() {
	counterOne := singleton.GetCounter()
	currentCount := counterOne.AddOne()
	fmt.Println(currentCount)

	counterTwo := singleton.GetCounter()
	currentCount = counterTwo.AddOne()
	fmt.Println(currentCount)
}
