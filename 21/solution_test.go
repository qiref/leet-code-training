package _1

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}
	_3 := &ListNode{Val: 4}
	_4 := &ListNode{Val: 5}

	_1.Next = _2
	_2.Next = _3
	_3.Next = _4

	_a := &ListNode{Val: 1}
	_b := &ListNode{Val: 3}
	_c := &ListNode{Val: 4}
	_d := &ListNode{Val: 5}

	_a.Next = _b
	_b.Next = _c
	_c.Next = _d

	fmt.Println(mergeTwoLists(_1, _a))
}
