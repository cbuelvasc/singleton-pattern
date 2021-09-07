package counter

import "testing"

func TestGetCounter(t *testing.T) {
	counterOne := GetCounter()
	if counterOne == nil {
		t.Error("A new connection object must have been made")
	}

	expectedCounter := counterOne
	currentCount := counterOne.AddOne()
	if currentCount != 1 || currentCount != counterOne.GetCount() {
		t.Errorf("After calling for the first time to count, the count must be 1 but it is %d\n", currentCount)
	}

	counterTwo := GetCounter()
	if counterTwo != expectedCounter {
		t.Error("Counter instances must be different")
	}

	currentCount = counterTwo.AddOne()
	if currentCount != 2 || currentCount != counterTwo.GetCount() {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was %d\n", currentCount)
	}

	counterThree := GetCounter()
	if counterThree != expectedCounter {
		t.Error("Counter instances must be different")
	}

	currentCount = counterThree.SubtractOne()
	if currentCount != 1 || currentCount != counterThree.GetCount() {
		t.Errorf("After calling 'SubtractOne' using the third counter, the current count must be 1 but was %d\n", currentCount)
	}
}
