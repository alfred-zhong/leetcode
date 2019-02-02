// https://leetcode.com/problems/merge-two-sorted-lists/description/
package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var r, l3 *ListNode

	for {
		var n *ListNode
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				n = l1
				l1 = l1.Next
			} else {
				n = l2
				l2 = l2.Next
			}
		} else if l1 == nil && l2 != nil {
			n = l2
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			n = l1
			l1 = l1.Next
		} else {
			break
		}

		if l3 == nil {
			l3 = n
			r = l3
		} else {
			l3.Next = n
			l3 = l3.Next
		}
	}

	return r
}
