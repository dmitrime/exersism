package summultiples

func MultipleSummer(divs ...int) func(int) int {
	return func(limit int) int {
		sum := 0
		for i := 1; i < limit; i++ {
			for _, d := range divs {
				if i%d == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
