package src

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


func rotateRight(head *ListNode, k int) *ListNode {
	count := 1
	headBackup := *head
	headPointer := &headBackup
	for{
		if head.Next != nil{
			count ++
			head = head.Next
		}else{
			break
		}
	}
	k = k%count
	for i:=0;i<k;i++{
		headPointer = rotateOnce(headPointer)
	}
	return headPointer
}


func rotateOnce(head *ListNode) *ListNode{
	//nodeHead := head
	var last, previous *ListNode
	previousHead :=head
	for{
		if head.Next != nil{
			previous =head
			head = head.Next
		}else{
			last = head
			previous.Next = nil
			break
		}
	}
	last.Next = previousHead
	return last
}

func PrintList(head *ListNode){
	for{
		fmt.Println(head.Val)
		if head.Next != nil{
			head =head.Next
		}else{
			break
		}
	}
}