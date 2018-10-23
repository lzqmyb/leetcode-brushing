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

func base_rotate(nums []int, k int) {
	if nums == nil || len(nums) == 0 || k%len(nums) == 0 {
		return
	}

	turns := k % len(nums)
	mid := len(nums) - turns
	reverse(nums, 0, mid-1)
	reverse(nums, mid, len(nums)-1)
	reverse(nums, 0, len(nums)-1)

}

func reverse(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start += 1
		end -= 1
	}
}

