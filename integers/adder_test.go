package integers

import "testing"

func TestAdder(t *testing.T) {
	sut := Add(2, 2)
	expected := 4

	if sut != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sut)
	}
}
