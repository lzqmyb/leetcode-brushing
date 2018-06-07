package eazy

//leetCode 21 题测试用列有问题， 不能通过https://leetcode-cn.com/problems/merge-two-sorted-lists/description/
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := &ListNode{0, nil}
	cl := l
	for l1 != nil || l2 != nil {
		if l1 != nil {
			cl.Next = l1
			cl = cl.Next
			l1 = l1.Next
		}

		if l2 != nil {
			cl.Next = l2
			cl = cl.Next
			l2 = l2.Next
		}
	}

	return l.Next
}
