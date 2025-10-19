package problem_0004_test

import (
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode/problem_0004"
)

func FindMedianSortedArrays_test(t *testing.T, nums1 []int, nums2 []int, expected float64) {
	result := FindMedianSortedArrays(nums1, nums2)
	if result != expected {
		t.Errorf("nums1 = %#v", nums1)
		t.Errorf("nums2 = %#v", nums2)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = FindMedianSortedArrays(nums1, nums2)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	FindMedianSortedArrays_test(t, []int{1, 3}, []int{2}, 2.0)
}

func TestExample2(t *testing.T) {
	FindMedianSortedArrays_test(t, []int{1, 2}, []int{3, 4}, 2.5)
}
