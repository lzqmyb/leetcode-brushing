package eazy

func FindMaxConsecutiveOnes(nums []int) int {
	var result int

	if len(nums) == 0 {
		panic("nums lenth is zero")
	}

	length := 0
	for _, v := range nums {
		if v == 1 {
			length++
			if result < length {
				result = length
			}
		} else {
			length = 0
		}
	}

	return result
}

func FindMaxConsecutiveOnesBetter(nums []int) int {
	if len(nums) == 0 {
		panic("error")
	}
	maxCount := 0
	count := 0
	length := len(nums)

	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			count++
		} else {
			if maxCount < count {
				maxCount = count
			}

			count = 0
		}
	}

	if maxCount < count {
		maxCount = count
	}

	return maxCount

}
