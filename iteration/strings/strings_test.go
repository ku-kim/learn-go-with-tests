package strings

import (
	"strings"
	"testing"
)

/**
https://pkg.go.dev/strings에 있는 라이브러리를 자유롭게 학습합니다.
*/

func TestContains_case_True(t *testing.T) {
	result := strings.Contains("abcde", "abc")

	if result != true {
		t.Errorf("테스트 실패")
	}
}

func TestContains_case_False(t *testing.T) {
	result := strings.Contains("abcde", "1")

	if result != false {
		t.Errorf("테스트 실패")
	}
}

func TestCount(t *testing.T) {
	result := strings.Count("abcdea", "a")
	expected := 2

	if result != expected {
		t.Errorf("테스트 실패")
	}
}
