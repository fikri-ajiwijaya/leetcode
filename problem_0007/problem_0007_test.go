package problem_0007_test

import (
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode.go/problem_0007"
)

func Reverse_test(t *testing.T, x int, expected int) {
	result := Reverse(x)
	if result != expected {
		t.Errorf("x = %#v", x)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = Reverse(x)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	Reverse_test(t, 123, 321)
}

func TestExample2(t *testing.T) {
	Reverse_test(t, -123, -321)
}

func TestExample3(t *testing.T) {
	Reverse_test(t, 120, 21)
}
