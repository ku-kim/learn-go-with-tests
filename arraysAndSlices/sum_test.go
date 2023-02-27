package arraysAndSlices

import "testing"

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}

	sut := Sum(numbers)

	if sut != 15 {
		t.Errorf("got %d want %d given, %v", sut, 15, numbers)
	}

}
