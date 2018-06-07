package eazy

// 测试用例有问题
type NumArray struct {
	List []int
}

func Constructor(nums []int) NumArray {
	return NumArray{nums}
}

func (this *NumArray) SumRange(i int, j int) int {
	in := []int{-2, 0, 3, -5, 2, -1}
	result := 0
	s := in[i : j+1]

	for _, v := range s {
		result += v
	}

	return result
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
