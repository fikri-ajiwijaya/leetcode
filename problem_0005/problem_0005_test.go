package problem_0005_test

import (
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode.go/problem_0005"
)

func LongestPalindrome_test(t *testing.T, s string, expected string) {
	result := LongestPalindrome(s)
	if result != expected {
		t.Errorf("s = %#v", s)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = LongestPalindrome(s)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	LongestPalindrome_test(t, "babad", "bab")
}

func TestExample2(t *testing.T) {
	LongestPalindrome_test(t, "cbbd", "bb")
}
