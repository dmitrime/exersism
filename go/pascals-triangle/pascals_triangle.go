package pascal

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	res := make([][]int, n)
	r := []int{1}
	res[0] = r
	for i := 1; i < n; i++ {
		t := r
		r = make([]int, i+1) // had a bug when using := here!
		r[0] = 1
		r[i] = 1
		for j := 1; j < i; j++ {
			r[j] = t[j-1] + t[j]
		}
		res[i] = r
	}
	return res
}
