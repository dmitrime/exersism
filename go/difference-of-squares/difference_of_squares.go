package diffsquares

func SquareOfSums(n int) int {
	sq := n * n
	return (sq + 2*sq*n + sq*sq) / 4
}

func SumOfSquares(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
