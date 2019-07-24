package eazy

import "strings"

func lengthOfLongestSubstring(s string) int {
	var sum,j = 0,0

	for i := 1; i < len(s); i++ {
		slice := s[j:i]
		curStr := s[i:i+1]
		if index := strings.Index(slice, curStr); index == -1 && sum < i-j+1{
			sum = i - j + 1
		} else {
			j = j + index + 1
		}
	}
	if (len(s) > 0 && sum == 0) {
		return 1
	}
	return sum
}
