package eazy

/**


Write a program to check whether a given number is an ugly number.

Ugly numbers are positive numbers whose prime factors only include 2, 3, 5.

Input: 6
Output: true
Explanation: 6 = 2 × 3

Input: 8
Output: true
Explanation: 8 = 2 × 2 × 2

Input: 14
Output: false
Explanation: 14 is not ugly since it includes another prime factor 7.

*/
func IsUgly(num int) bool {
	result, ok := Aliquot(num, 2)
	if ok {
		return true
	}

	//fmt.Println("result: ", result)
	result, ok = Aliquot(result, 3)
	if ok {
		return true
	}

	//fmt.Println("result: ", result)
	result, ok = Aliquot(result, 5)
	if ok {
		return true
	}

	return false
}

func Aliquot(num, base int) (int, bool) {
	if num == 0 {
		return 0, false
	}
	for {
		remainder, yushu := num/base, num%base

		if yushu == 1 && remainder == 0 {
			return 1, true
		}

		if yushu == 0 {
			//fmt.Printf("%v , %v, %v \n", remainder, yushu, base)
			num = remainder
		} else {
			return num, false
		}
	}
}

func isUglyBetter(num int) bool {
	if num < 1 {
		return false
	}

	for num%2 == 0 {
		num /= 2
	}

	for num%3 == 0 {
		num /= 3
	}

	for num%5 == 0 {
		num /= 5
	}

	return num == 1
}
