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

func TestExample4(t *testing.T) {
	Reverse_test(t, 1534236469, 0)
}

func TestExample5(t *testing.T) {
	Reverse_test(t, 0, 0)
}

func TestExample6(t *testing.T) {
	Reverse_test(t, 7463847412, 2147483647)
}

func TestExample7(t *testing.T) {
	Reverse_test(t, 8463847412, 0)
}

func TestExample8(t *testing.T) {
	Reverse_test(t, -8463847412, -2147483648)
}

func TestExample9(t *testing.T) {
	Reverse_test(t, -9463847412, 0)
}
