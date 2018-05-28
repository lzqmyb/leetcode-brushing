package array

import "fmt"

func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return []int{-1,-1}
	}

	length := len(nums)
	index := 0
	uniqueList := map[int]bool{}
con1:
	for i:=0 ; i< length - 1; i++ {
		_, ok := uniqueList[nums[index]]
		if ok {
			//fmt.Println(i,"==", "continue")
			index++
			continue con1
		}
		for j:= index + 1; j < length; j++{

			//fmt.Println("j == ", j, " result= ", nums[index] + nums[j])
			if sum := nums[index] + nums[j]; sum == target {
				return []int{index, j}
			}
		}
		uniqueList[nums[index]] = true
		index++
	}
	return []int{-1, -1}
}

func twoSumBetter(nums []int, target int) interface{} {
	result := make([]int, 2)
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if index, ok := m[target - nums[i]]; ok {
			result[1] = i
			result[0] = index
			return result
		}
		m[nums[i]] = i
	}
	return result
}

func main() {
	arr := []int{3,2,3}
	fmt.Println(twoSum(arr, 6))
	fmt.Println(twoSumBetter(arr, 6))
}
