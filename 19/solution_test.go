package _9

import (
	"fmt"
	"testing"
)

func Test_removeNthFromEnd(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}

	_1.Next = _2

	fmt.Println(removeNthFromEnd(_1, 1))
}

func Test_removeNthFromEnd1(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}

	_1.Next = _2

	fmt.Println(removeNthFromEnd(_1, 2))
}
