package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	if limit <= 1 {
		return 0
	}
	seen := make([]bool, limit)
	total := 0
	for _, d := range divisors {
		if d <= 0 {
			continue
		}
		for n := d; n < limit ; n += d {
			if !seen[n] {
				seen[n] = true
				total += n
			}
		}
	}
	return total
}
