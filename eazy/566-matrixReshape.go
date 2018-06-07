package eazy

func matrixReshape(nums [][]int, r int, c int) [][]int {
	rOld, cOld := len(nums), len(nums[0])
	if r*c != rOld*cOld {
		return nums
	}

	result := [][]int{}
	item := []int{}
	rindex := 0
	cindex := 0

	for i := 0; i < rOld; i++ {
		for y := 0; y < cOld; y++ {
			//fmt.Printf("i: %v, y: %v \n", i, y)
			item = append(item, nums[i][y])
			if cindex++; cindex == c {
				result = append(result, item)
				item = []int{}
				cindex = 0
				rindex++
			}
		}
	}

	return result
}

func matrixReshapeBetter(nums [][]int, r int, c int) [][]int {
	rows := len(nums)
	if rows == 0 {
		return nil
	}

	cols := len(nums[0])
	if cols == 0 {
		return nil
	}

	if rows*cols != r*c {
		return nums
	}

	ans := make([][]int, 0, r)
	for i := 0; i < r; i++ {
		coli := make([]int, 0, c)
		for j := 0; j < c; j++ {
			coli = append(coli, nums[(i*c+j)/cols][(i*c+j)%cols])
		}

		ans = append(ans, coli)
	}

	return ans
}
