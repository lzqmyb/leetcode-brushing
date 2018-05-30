package array

type ListNode struct {
	Val int
	Next *ListNode
}

func GenerateListNode(in []int) *ListNode {
	result := ListNode{}
	current := &result

	lastIndex := len(in) - 1
	for index, v := range in {
		current.Val = v
		if index != lastIndex {
			next := ListNode{}
			current.Next = &next
			current = &next
		}
	}

	return &result
}
