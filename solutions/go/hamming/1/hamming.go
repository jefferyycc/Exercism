package hamming

import "errors"

var ErrDifferentLength = errors.New("DNA strands are different lengths")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, ErrDifferentLength
	}

	distance := 0
	for i := range a {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
