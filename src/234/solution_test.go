package _34

import (
	"fmt"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}
	_3 := &ListNode{Val: 2}
	_4 := &ListNode{Val: 1}

	_1.Next = _2
	_2.Next = _3
	_3.Next = _4
	fmt.Println(isPalindrome1(_1))
}

func Test_isPalindrome1(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}
	_3 := &ListNode{Val: 2}
	_4 := &ListNode{Val: 1}

	_1.Next = _2
	_2.Next = _3
	_3.Next = _4
	fmt.Println(isPalindrome2(_1))
}

func Test_isPalindrome2(t *testing.T) {
	_1 := &ListNode{Val: 1}
	_2 := &ListNode{Val: 2}
	_21 := &ListNode{Val: 3}
	_3 := &ListNode{Val: 2}
	_4 := &ListNode{Val: 1}

	_1.Next = _2
	_2.Next = _21
	_21.Next = _3
	_3.Next = _4
	fmt.Println(isPalindrome2(_1))
}

func Test_isPalindrome3(t *testing.T) {
	_1 := &ListNode{Val: 2}
	_2 := &ListNode{Val: 2}
	_4 := &ListNode{Val: 1}

	_1.Next = _2
	_2.Next = _4

	fmt.Println(isPalindrome2(_1))
}
