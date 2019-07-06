package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	over := false
	head := l1
	for l1 != nil && l2 != nil {
		if over {
			l1.Val += l2.Val + 1
		} else {
			l1.Val += l2.Val
		}

		if l1.Val >= 10 {
			l1.Val -= 10
			over = true
		}else{
			over = false
		}
		if l1.Next == nil && l2.Next != nil{
			l1.Next = &ListNode{}
		}

		if l1.Next != nil && l2.Next == nil{
			l2.Next = &ListNode{}
		}

		if over && l1.Next == nil && l2.Next == nil{
			l1.Next = &ListNode{1,nil}
		}

		l1 = l1.Next
		l2 = l2.Next
	}

	return head
}
