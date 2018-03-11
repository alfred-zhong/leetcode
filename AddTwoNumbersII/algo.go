// https://leetcode.com/problems/add-two-numbers-ii/description/
package algo

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	for l1 != nil {
		n := &ListNode{
			Val:  l1.Val,
			Next: l3,
		}
		l3 = n

		l1 = l1.Next
	}

	var l4 *ListNode
	for l2 != nil {
		n := &ListNode{
			Val:  l2.Val,
			Next: l4,
		}
		l4 = n

		l2 = l2.Next
	}

	var l5 *ListNode

	carry := false
	for {
		if l3 != nil || l4 != nil || carry {
			n := new(ListNode)

			if l3 != nil {
				n.Val += l3.Val
				l3 = l3.Next
			}
			if l4 != nil {
				n.Val += l4.Val
				l4 = l4.Next
			}
			if carry {
				n.Val++
				carry = false
			}

			if n.Val >= 10 {
				carry = true
			}
			n.Val %= 10

			n.Next = l5
			l5 = n
		} else {
			break
		}
	}

	return l5
}
