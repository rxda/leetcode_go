package src


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	pointer := head
	start := pointer
	lastVal := head.Val
	for{
		if pointer.Next != nil{
			temp := pointer
			flag := false
			for lastVal == temp.Next.Val{
				flag = true
				temp = temp.Next
			}
			if flag{
				pointer = temp.Next
				lastVal = pointer.Val
				start = pointer
			}else{
				pointer = pointer.Next
			}
		}else{
			break
		}
	}
	return start
}


func createList(input []int) *ListNode{
	result := &ListNode{}
	result2 := result
	length := len(input)
	for k,v := range input{
		result2.Val = v
		if k != length -1{
			result2.Next = &ListNode{}
			result2 = result2.Next
		}
	}
	return result
}