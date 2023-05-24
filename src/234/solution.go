package _34

/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
 */

// 给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。
//输入：head = [1,2,2,1]
//输出：true
//输入：head = [1,2]
//输出：false
//提示：
//
//链表中节点数目在范围[1, 105] 内
//0 <= Node.val <= 9
//
//进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

type ListNode struct {
	Val  int
	Next *ListNode
}

// isPalindrome 反转最简单的就是借助数组，把数据放到数组中，然后反序遍历，跟之前的链表元素对比
func isPalindrome(head *ListNode) bool {
	reserveArr := make([]int, 0)
	origin := head
	for head.Next != nil {
		reserveArr = append(reserveArr, head.Val)
		head = head.Next
	}
	reserveArr = append(reserveArr, head.Val)

	for i := len(reserveArr) - 1; i >= 0; i-- {
		if reserveArr[i] != origin.Val {
			return false
		}
		origin = origin.Next
	}
	return true
}

func isPalindrome1(head *ListNode) bool {
	reserveArr := make([]int, 0)
	origin := head
	for ; head != nil; head = head.Next {
		reserveArr = append(reserveArr, head.Val)
	}

	for i := len(reserveArr) - 1; i >= 0; i-- {
		if reserveArr[i] != origin.Val {
			return false
		}
		origin = origin.Next
	}
	return true
}

// isPalindrome2 借助快慢指针，快指针每次移动两个元素，当 fast 移动到最后时，慢指针恰好在链表中间位置，反转前半部分链表，然后跟后半部分链表对比
// 问题：如何反转前半部分链表？ 可以借助链表反转（ 206 ）的结果，链表反转时会把链表分成两部分，恰好可以对比
// 问题：奇数和偶数不同，会导致链表分成的两部分不对称，如何对比？如果是奇数，中间的元素不用对比，偶数直接对比；
//      使用快指针时，当 fast = nil 时，说明是偶数，当fast != nil 时，是奇数，此时需要把 后半部分链表向后移动一个元素，过滤中间元素
//
func isPalindrome2(head *ListNode) bool {
	if head == nil {
		return true
	}
	if head.Next == nil {
		return true
	}
	var begin *ListNode
	mid := head
	end := head.Next

	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		mid.Next = begin
		begin = mid
		mid = end
		end = end.Next
	}
	//mid.Next = begin

	if fast != nil {
		mid = mid.Next
	}

	for begin != nil && mid != nil {
		if begin.Val != mid.Val {
			return false
		}
		mid = mid.Next
		begin = begin.Next
	}
	return true
}
