package eazy

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	result = &ListNode{0, nil}
	if l1 == nil && l2 == nil {
		return result
	}
	return getNumbers(l1, l2, 0, result, result)
}

func getNumbers(l1 *ListNode, l2 *ListNode, carry int, current *ListNode, result *ListNode) *ListNode {

	sum := 0
	switch {
	case l1 != nil && l2 != nil:
		sum = l1.Val + l2.Val
	case l1 != nil:
		sum = l1.Val
	case l2 != nil:
		sum = l2.Val
	}
	current.Val = (sum + carry) % 10
	carry = (sum + carry) / 10
	if (l1 == nil || l1.Next == nil) && (l2 == nil || l2.Next == nil) && carry == 0 {
		return result
	}
	current.Next = &ListNode{}
	if l1 != nil {
		l1 = l1.Next
	}
	if l2 != nil {
		l2 = l2.Next
	}
	return getNumbers(l1, l2, carry, current.Next, result)
}

func AddTwoNumbersBetter(l1 *ListNode, l2 *ListNode) (result *ListNode) {
	carry := 0
	sum := 0

	p := &ListNode{
		Val:  0,
		Next: nil,
	}
	head := p

	for l1 != nil || l2 != nil || carry != 0 {
		sum = carry

		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		p.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		carry = sum / 10
		p = p.Next
	}
	return head.Next
}

//copy
func addTwoNumbersBetterBetter(l1 *ListNode, l2 *ListNode) *ListNode {
	var head = &ListNode{}
	var tail *ListNode
	var length1, length2 = 0, 0
	for tail = l1; tail != nil; tail = tail.Next {
		length1++
	}
	for tail = l2; tail != nil; tail = tail.Next {
		length2++
	}
	var length int
	if length1 > length2 {
		length = length1
	} else {
		length = length2
	}
	var tmp int
	tail = head
	var left, right = l1, l2
	for i := length; i > 0; i-- {
		if left != nil {
			tmp = tmp + left.Val
			left = left.Next
		}
		if right != nil {
			tmp = tmp + right.Val
			right = right.Next
		}
		tail.Next = &ListNode{tmp % 10, nil}
		tmp = tmp / 10
		tail = tail.Next
	}
	if tmp > 0 {
		tail.Next = &ListNode{tmp, nil}
	}
	return head.Next
}
