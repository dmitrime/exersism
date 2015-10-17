package cryptosquare

import (
	"math"
	"strings"
)

const TestVersion = 1

func Encode(s string) string {
	t := strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, strings.ToLower(s))
	cols := int(math.Ceil(math.Sqrt(float64(len(t)))))

	var res string
	for i := 0; i < cols; i++ {
		for j := i; j < len(t); j += cols {
			res += string(t[j])
		}
		if i != cols-1 {
			res += " "
		}
	}
	return res
}
