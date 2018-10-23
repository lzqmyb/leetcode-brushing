package utils

func ListSwap(nums []int, i, j int)  {
	z := nums[i]
	nums[i] = nums[j]
	nums[j] = z
}
