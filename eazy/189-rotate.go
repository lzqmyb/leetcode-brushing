package eazy

import "fmt"

func  rotate(nums []int, k int) []int {
	length := len(nums)
	// if (k + 1 > length) {
	// 	length = k + 1
	// }
	k = k % length

	result := make([]int, length)
	copy(result[:k], nums[length - k:])
	copy(result[k:], nums[:length - k])
	copy(nums, result)

	fmt.Println(nums)
	return nums
}

