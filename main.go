package main

import (
	"leetcode/array"
	"fmt"
)

func main()  {
	//l1 := array.GenerateListNode([]int{1})
	//l2 := array.GenerateListNode([]int{9,9})
	//result := array.AddTwoNumbers(l1,l2)
	//for result != nil {
	//	fmt.Println(result.Val)
	//	result = result.Next
	//}
	result := array.SingleNumber([]int{2,2,1})
	result = array.SingleNumberBetter([]int{4,1,2,1,2})
	fmt.Println(result)
}

