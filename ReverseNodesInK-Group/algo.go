// https://leetcode.com/problems/reverse-nodes-in-k-group/
package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var curr *ListNode = head

	var count int = 0
	for curr != nil && count != k {
		curr = curr.Next
		count++
	}

	if count == k {
		curr = reverseKGroup(curr, k)

		for count > 0 {
			temp := head.Next
			head.Next = curr
			curr = head
			head = temp

			count--
		}
		head = curr
	}
	return head
}
