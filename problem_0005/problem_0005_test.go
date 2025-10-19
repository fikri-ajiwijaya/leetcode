package problem_0005_test

import (
	"slices"
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode/problem_0005"
)

func LongestPalindrome_test(t *testing.T, s string, expected []string) {
	result := LongestPalindrome(s)
	if !slices.Contains(expected, result) {
		t.Errorf("s = %#v", s)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = LongestPalindrome(s)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	LongestPalindrome_test(t, "babad", []string{"bab", "aba"})
}

func TestExample2(t *testing.T) {
	LongestPalindrome_test(t, "cbbd", []string{"bb"})
}

func TestExample3(t *testing.T) {
	LongestPalindrome_test(t, "bb", []string{"bb"})
}

func TestExample4(t *testing.T) {
	LongestPalindrome_test(t, "aacabdkacaa", []string{"aca"})
}
