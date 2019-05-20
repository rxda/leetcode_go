package src

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode) *ListNode{
	cloneList := *head
	length := 1
	nodeSlice :=[]*ListNode{}
	nodeSlice = append(nodeSlice, &cloneList)
	temp := head
	for temp.Next!=nil{
		length++
		nodeSlice = append(nodeSlice, temp.Next)
		temp = temp.Next
	}
	n:=2
	nodeSlice[length-n-1].Next = nodeSlice[length-n+1]
	return nodeSlice[0]
}
