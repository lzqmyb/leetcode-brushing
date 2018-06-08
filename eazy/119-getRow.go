package eazy

/**
Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
Note that the row index starts from 0.
Input: 3
Output: [1,3,3,1]
*/
func GetRow(rowIndex int) []int {
	start := []int{1}

	for i := 0; i < rowIndex; i++ {
		result := []int{}
		for y := 0; y < i+2; y++ {
			if y == 0 {
				result = append(result, start[y])
			} else if y-1 == i {
				result = append(result, start[y-1])
			} else {
				result = append(result, start[y]+start[y-1])
			}
		}
		start = result
	}

	return start
}

var nums = [32][]int{
	{1},
	{1, 1},
}

func getRowBetter(rowIndex int) []int {
	if len(nums[rowIndex]) != 0 {
		return nums[rowIndex]
	}

	for i := 1; i <= rowIndex; i++ {
		nums[i] = make([]int, i+1)
		nums[i][0], nums[i][i] = 1, 1
		for j := 1; j < i; j++ {
			nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
		}
	}
	return nums[rowIndex]

	//for i := 7; i <= rowIndex; i++ {
	//	nums[i] = make([]int, i+1)
	//	nums[i][0], nums[i][i] = 1, 1
	//	for j := 1; j < i; j++ {
	//		nums[i][j] = nums[i-1][j-1] + nums[i-1][j]
	//	}
	//}
	//return nums[rowIndex]

}
