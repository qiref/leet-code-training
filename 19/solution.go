package _9

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//示例 1：
//
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
//
//示例 2：
//输入：head = [1], n = 1
//输出：[]
//
//示例 3：
//
//输入：head = [1,2], n = 1
//输出：[1]

type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd
// 思路：倒数第 n 个，如果 n = 1 那么 n.next = nil ; 如果 n = 2 那么 n.next.next = nil ;
// 所以遍历链表时，加一个for循环，往后循环 n 次，直到 n.next = nil , 则找到倒数第 n 个元素；
// 问题1: 需要有一个指针保留head，所以遍历链表前，origin := head
// 问题2: 边界问题，prev.Next = prev.Next.Next，如果  prev.Next = nil ，此时会报空指针；
// 问题3: 边界问题，当要删除的节点为头节点时，prev.Next = prev.Next.Next 就不对了；需要兼容这种情况，还是回到之前的思路：
//        倒数第 n 个，如果 n = 1 那么 n.next = nil ; 如果 n = 2 那么 n.next.next = nil；正常情况下，i = n 时，满足条件;
//        当删除的是头节点时，会出现身情况呢？ 1->2->nil 删除倒数第2个，i = 1 时，2.next = nil 成立，此时 i < n;
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil || n <= 0 {
		return head
	}
	if head.Next == nil {
		return nil
	}
	origin := head
	for ; head != nil; head = head.Next {
		prev := head
		tmp := head
		for i := 0; i <= n; i++ {
			if tmp.Next == nil {
				if i < n {
					// 如果要删除的是头部节点
					result := origin.Next
					origin.Next = nil
					return result
				}
				// 找到第k个节点
				if prev.Next == nil {
					prev.Next = nil
					return origin
				}
				prev.Next = prev.Next.Next
				return origin
			}
			tmp = tmp.Next
		}
	}
	return origin
}
