package armstrongnumbers

func IsNumber(n int) bool {
	if n < 0 {
		return false
	}

	digits := 1
	for remaining := n; remaining >= 10; remaining /= 10 {
		digits++
	}

	sum := 0
	for remaining := n; remaining > 0; remaining /= 10 {
		digit := remaining % 10
		power := 1
		for i := 0; i < digits; i++ {
			power *= digit
		}
		sum += power
	}
	return sum == n
}
