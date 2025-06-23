package problem_0006_test

import (
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode.go/problem_0006"
)

func Convert_test(t *testing.T, s string, numRows int, expected string) {
	result := Convert(s, numRows)
	if result != expected {
		t.Errorf("s = %#v", s)
		t.Errorf("numRows = %#v", numRows)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = Convert(s, numRows)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	Convert_test(t, "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR")
}

func TestExample2(t *testing.T) {
	Convert_test(t, "PAYPALISHIRING", 4, "PINALSIGYAHRPI")
}

func TestExample3(t *testing.T) {
	Convert_test(t, "A", 1, "A")
}
