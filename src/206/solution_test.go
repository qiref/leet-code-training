package _06

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}
	_3 := &ListNode{Val: 2}
	_4 := &ListNode{Val: 1}

	_1.Next = _2
	_2.Next = _3
	_3.Next = _4
	fmt.Println(reverseList(_1))
}
