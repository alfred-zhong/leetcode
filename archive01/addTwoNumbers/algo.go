package addTwoNumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ln := new(ListNode)
	carry := 0
	var r *ListNode
	for n1, n2 := l1, l2; n1 != nil || n2 != nil || carry > 0; {
		if r == nil {
			r = ln
		} else {
			r.Next = new(ListNode)
			r = r.Next
		}

		sum := carry
		if n1 != nil {
			sum += n1.Val
			n1 = n1.Next
		}
		if n2 != nil {
			sum += n2.Val
			n2 = n2.Next
		}

		carry = sum / 10
		r.Val = sum % 10
	}
	return ln
}
