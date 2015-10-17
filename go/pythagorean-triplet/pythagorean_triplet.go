package pythagorean

type Triplet [3]int

func Range(min, max int) []Triplet {
	res := []Triplet{}
	for i := min; i <= max; i++ {
		for j := i + 1; j <= max; j++ {
			for k := j + 1; k <= max; k++ {
				if i*i+j*j == k*k {
					res = append(res, Triplet{i, j, k})
				}
			}
		}
	}
	return res
}

func Sum(p int) []Triplet {
	res := []Triplet{}
	tri := Range(1, p-1)
	for _, t := range tri {
		if t[0]+t[1]+t[2] == p {
			res = append(res, t)
		}
	}
	return res
}
