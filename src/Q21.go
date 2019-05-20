package src

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	dummy := ListNode{
		Val:  -1,
		Next: nil,
	}
	cur := &dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
		}
		cur = cur.Next
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return dummy.Next

}

func MergeKLists(lists []*ListNode) *ListNode {
	i:=0
	j:= len(lists)-1
	for i<j{
		lists[i+1] = MergeTwoLists(lists[i],lists[i+1])
		lists[j+1] = MergeTwoLists(lists[j-1],lists[j])
		i++
		j--
	}
	return lists[len(lists)-1]


}

func SwapPairs(head *ListNode) *ListNode {
	cur := head
	for cur!=nil &&cur.Next!=nil{
		temp :=cur.Val
		cur.Val =cur.Next.Val
		cur.Next.Val =temp
		cur =cur.Next.Next
	}
	return head
}

