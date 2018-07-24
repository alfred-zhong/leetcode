// https://leetcode.com/problems/reverse-linked-list/description/
package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		temp := head.Next
		head.Next = prev
		prev = head
		head = temp
	}
	return prev
}
