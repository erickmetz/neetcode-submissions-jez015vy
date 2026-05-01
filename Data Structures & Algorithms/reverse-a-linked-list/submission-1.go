func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head

	for cur != nil {
		oldNext := cur.Next
		cur.Next = prev
		prev = cur
		cur = oldNext
	}

	return prev
}
