package _06

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//输入：head = [1,2,3,4,5]
//输出：[5,4,3,2,1]

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 链表反转，用3个指针，第一个用于记录反转后的结果，mid 负责反转，mid.next = begin
// mid.next 指向 begin 之后，mid 与后面的节点就断了，所以需要 end 节点，保存原始链表的头部
// 每次循环，指针都向后移动，直到 end = nil
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var begin *ListNode = nil
	mid := head
	end := head.Next

	for {
		mid.Next = begin
		if end == nil {
			break
		}
		begin = mid
		mid = end
		end = end.Next

	}
	return mid
}
