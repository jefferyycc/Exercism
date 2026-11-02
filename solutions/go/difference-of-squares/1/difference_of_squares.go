package differenceofsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := range n {
		sum += i + 1
	}
	return sum * sum
}

func SumOfSquares(n int) int {
	sum := 0
	for i := range n {
		sum += (i + 1) * (i + 1)
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
