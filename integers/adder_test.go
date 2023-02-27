package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sut := Add(2, 2)
	expected := 4

	if sut != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sut)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
