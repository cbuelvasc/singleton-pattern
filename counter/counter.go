package singleton

import "fmt"

type counter struct {
	accumulator int
}

var instance *counter

func GetCounter() *counter {
	if instance == nil {
		fmt.Println("Creating counter instance")
		instance = &counter{}
	} else {
		fmt.Println("Counter instance already created")
	}
	return instance
}

func (s *counter) AddOne() int {
	s.accumulator++
	return s.accumulator
}

func (s *counter) SubtractOne() int {
	s.accumulator--
	return s.accumulator
}

func (s *counter) GetCount() int {
	return s.accumulator
}
