package src

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	duMap := make(map[int]bool)
	start := head
	cursor := head
	duMap[head.Val] =true
	for cursor.Next != nil{
		if _,ok := duMap[cursor.Next.Val];ok{
			if cursor.Next.Next != nil{
				cursor.Next =cursor.Next.Next
			}else{
				cursor.Next = nil
				break
			}
		}else{
			duMap[cursor.Next.Val] =true
			cursor = cursor.Next
		}

	}
	return start

	//1,2,1,4,5


}