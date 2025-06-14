package problem_0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func Slice_to_List(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	return &ListNode{vals[0], Slice_to_List(vals[1:])}
}

func List_to_Slice(l *ListNode) []int {
	_l := []int{}
	for node := l; node != nil; node = node.Next {
		_l = append(_l, node.Val)
	}
	return _l
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbers(l1, l2)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	_l1 := List_to_Slice(l1)
	_l2 := List_to_Slice(l2)
	_l3 := []int{}
	carry := 0
	for i := 0; i < len(_l1) || i < len(_l2); i++ {
		if i < len(_l1) {
			carry += _l1[i]
		}
		if i < len(_l2) {
			carry += _l2[i]
		}
		_l3 = append(_l3, carry%10)
		carry /= 10
	}
	if carry != 0 {
		_l3 = append(_l3, carry)
	}
	l3 := Slice_to_List(_l3)
	return l3
}
