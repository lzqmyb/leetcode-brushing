package eazy

func isOneBitCharacter(bits []int) bool {
	if len(bits) == 1 {
		return true
	}
	sum := 0

	for i := len(bits) - 2; i >= 0; i-- {
		if bits[i] == 0 {
			break
		} else {
			sum++
		}
	}

	if sum%2 == 0 {
		return true
	}

	return false
}
