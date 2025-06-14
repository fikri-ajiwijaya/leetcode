package problem_0002_test

import (
	"reflect"
	"testing"

	. "github.com/fikri-ajiwijaya/leetcode.go/problem_0002"
)

func List_test(t *testing.T, l []int) {
	result := List_to_Slice(Slice_to_List(l))
	if !reflect.DeepEqual(l, result) {
		t.Errorf("l = %#v", l)
		t.Errorf("Expected: l == List_to_Slice(Slice_to_List(l))")
		t.Errorf("Found: %#v", result)
	}
}

func AddTwoNumbers_test(t *testing.T, l1 []int, l2 []int, expected []int) {
	l1_ := Slice_to_List(l1)
	l2_ := Slice_to_List(l2)
	result_ := AddTwoNumbers(l1_, l2_)
	result := List_to_Slice(result_)
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("l1 = %#v", l1)
		t.Errorf("l2 = %#v", l2)
		t.Errorf("expected = %#v", expected)
		t.Errorf("result = AddTwoNumbers(l1, l2)")
		t.Errorf("result is %#v", result)
	}
}

func TestListv1(t *testing.T) {
	List_test(t, []int{2, 4, 3})
}

func TestListv2(t *testing.T) {
	List_test(t, []int{5, 6, 4})
}

func TestListv3(t *testing.T) {
	List_test(t, []int{7, 0, 8})
}

func TestListv4(t *testing.T) {
	List_test(t, []int{0})
}

func TestListv5(t *testing.T) {
	List_test(t, []int{9, 9, 9, 9, 9, 9, 9})
}

func TestListv6(t *testing.T) {
	List_test(t, []int{9, 9, 9, 9})
}

func TestListv7(t *testing.T) {
	List_test(t, []int{8, 9, 9, 9, 0, 0, 0, 1})
}

func TestExample1(t *testing.T) {
	AddTwoNumbers_test(t, []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8})
}

func TestExample2(t *testing.T) {
	AddTwoNumbers_test(t, []int{0}, []int{0}, []int{0})
}

func TestExample3(t *testing.T) {
	AddTwoNumbers_test(t, []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1})
}
