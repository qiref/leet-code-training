package _41

//给你一个链表的头节点 head ，判断链表中是否有环。
//如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。
//为了表示给定链表中的环，评测系统内部使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。
//注意：pos 不作为参数进行传递 。仅仅是为了标识链表的实际情况。
//如果链表中存在环 ，则返回 true 。 否则，返回 fal
//
//输入：head = [3,2,0,-4], pos = 1
//输出：true
//解释：链表中有一个环，其尾部连接到第二个节点。se 。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle
// 思路：如果有环，某个节点会出现多次，所以用 map 保存节点，当发现 map 中已存在节点时，则说明节点出现多次，链表有环
func hasCycle(head *ListNode) bool {
	mapVal := make(map[*ListNode]int)
	for ; head != nil; head = head.Next {
		if _, ok := mapVal[head]; ok {
			return true
		} else {
			mapVal[head] = 1
		}
	}
	return false
}

// hasCycle1
// 思路：使用快慢指针，当链表有环时，快慢指针都会进入环中循环，最终快慢指针一定会相遇
// 问题：如何证明，快慢指针一定会相遇；
func hasCycle1(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	low := head
	fast := head.Next.Next
	result := false
	for {
		if low == nil || fast == nil || fast.Next == nil {
			result = false
			break
		}
		if low == fast {
			result = true
			break
		}
		low = low.Next
		fast = fast.Next.Next
	}
	return result
}
