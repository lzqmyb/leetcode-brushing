package eazy

func SingleNumber(nums []int) int {
	m := map[int]bool{}

	for _, v := range nums {
		if _, ok := m[v]; ok {
			delete(m, v)
		} else {
			m[v] = true
		}
	}

	for k, v := range m {
		if v {
			return k
		} else {
			return -1
		}
	}
	return -1
}

func SingleNumberBetter(nums []int) int {
	if len(nums) == 0 {
		panic("no Single nums")
	}

	result := 0

	for _, v := range nums {
		result ^= v
	}

	return result
}
