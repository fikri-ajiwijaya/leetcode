package problem_0001_test

import (
	"reflect"
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode/problem_0001"
)

func TwoSum_test(t *testing.T, nums []int, target int, expected []int) {
	result := TwoSum(nums, target)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("nums = %#v", nums)
		t.Errorf("target = %#v", target)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = TwoSum(nums, target)")
		t.Errorf("result is %#v", result)
	}
}

func TestExample1(t *testing.T) {
	TwoSum_test(t, []int{2, 7, 11, 15}, 9, []int{0, 1})
}

func TestExample2(t *testing.T) {
	TwoSum_test(t, []int{3, 2, 4}, 6, []int{1, 2})
}

func TestExample3(t *testing.T) {
	TwoSum_test(t, []int{3, 3}, 6, []int{0, 1})
}
