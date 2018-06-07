package eazy

func RemoveElements(head *ListNode, val int) *ListNode {

	if head == nil {
		return nil
	}

	result := &ListNode{}
	for head != nil && head.Val == val {
		head = head.Next
	}
	result = head

	if result == nil {
		return nil
	}

	var beforeNode = head
	head = head.Next

	for head != nil {
		if head.Val == val {
			beforeNode.Next = head.Next
		} else {
			beforeNode = beforeNode.Next
		}
		if head == nil {
			return result
		}
		head = head.Next
	}
	return result
}

func rewrite(head *ListNode, val int) *ListNode {
	var current_addr *ListNode
	fakeNode := ListNode{
		0,
		head,
	}
	current_addr = &fakeNode

	for {
		if current_addr.Next != nil {
			nextNode := current_addr.Next
			if nextNode.Val == val {
				current_addr.Next = nextNode.Next
			} else {
				current_addr = current_addr.Next
			}

		} else {
			break
		}
	}

	return fakeNode.Next
}

func RemoveElementsBetter(head *ListNode, val int) *ListNode {
	var current_addr *ListNode
	fake_head := ListNode{
		Val:  0,
		Next: head,
	}
	current_addr = &fake_head
	for {
		if (*current_addr).Next != nil {
			next_node := *(*current_addr).Next
			if next_node.Val == val {
				(*current_addr).Next = next_node.Next
			} else {
				current_addr = (*current_addr).Next
			}
		} else {
			break
		}
	}
	return fake_head.Next
}
