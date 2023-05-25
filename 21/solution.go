package _1

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
//示例 1：
//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
//
//示例 2：
//输入：l1 = [], l2 = []
//输出：[]
//
//示例 3：
//输入：l1 = [], l2 = [0]
//输出：[0]

type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists
// 链表合并思路：创建一个新链表接收合并后的结果，如果 list1.val < list2.val then result.Next = list1 ;
// 注意此时 list1 向后移动，result 也需要向后移动;
// 所以最后返回的不是 result，需要提前保留 result.head
// 如果 list1.len != list2.len 最后一定会某个链表还有剩余，result 还要加上剩余的节点；
// tips: 菜逼写法
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	result := &ListNode{}
	resultHead := result
	if list1.Val > list2.Val {
		result.Val = list2.Val
		list2 = list2.Next
	} else {
		result.Val = list1.Val
		list1 = list1.Next
	}
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			tmp := list2
			result.Next = tmp
			list2 = list2.Next
			tmp.Next = nil
			result = result.Next
		} else {
			tmp := list1
			result.Next = tmp
			list1 = list1.Next
			tmp.Next = nil
			result = result.Next
		}
	}
	if list1 == nil {
		result.Next = list2
	}
	if list2 == nil {
		result.Next = list1
	}

	return resultHead
}

// mergeTwoLists1
// tips: 优雅的递归写法
// f(x) = x*f(x-1)
// l1.next = f(l1.next,l2) 怎么理解？ => 合并后的链表就是 l1(head) + (f1.next,l2) 把后面的步骤当成整体看待
// 当两个链表不等长时，最后返回的部分就是递归的最后一步，然后开始合并
func mergeTwoLists1(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists1(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists1(list2.Next, list1)
		return list2
	}
}
