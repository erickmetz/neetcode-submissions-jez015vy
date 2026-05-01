func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for {
		if cur == nil {
			break
		} else if cur.Next == nil {
			cur.Next = prev
			head = cur
			break
		}

		oldNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = oldNext	
	}


	return head
}
