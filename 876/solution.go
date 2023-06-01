package _76

//给你单链表的头结点 head ，请你找出并返回链表的中间结点。
//
//如果有两个中间结点，则返回第二个中间结点。
//
//示例 1：
//输入：head = [1,2,3,4,5]
//输出：[3,4,5]
//解释：链表只有一个中间结点，值为 3 。
//
//示例 2：
//输入：head = [1,2,3,4,5,6]
//输出：[4,5,6]
//解释：该链表有两个中间结点，值分别为 3 和 4 ，返回第二个结点。

type ListNode struct {
	Val  int
	Next *ListNode
}

// middleNode
// 思路，快慢指针，当快指针到最后时，slow 刚好到中间节点；
// 问题1: 边界问题，当 head 为空或者 head.next 为空；
// 问题2: 链表的奇数个和偶数个情况不同；
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil || head == nil {
		return head
	}
	slow := head
	fast := head
	for ; fast != nil; fast = fast.Next.Next {
		if fast.Next == nil {
			return slow
		}
		slow = slow.Next
	}
	if fast == nil {
		return slow
	}
	return nil
}
