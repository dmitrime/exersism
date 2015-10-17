package grains

import "errors"

func Exp(n int) uint64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	s := Exp(n / 2)
	if n%2 == 1 {
		return s * s * 2
	} else {
		return s * s
	}
}

func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("bad value")
	}
	return Exp(n - 1), nil
}

func Total() uint64 {
	var sum uint64 = 0
	for i := 1; i <= 64; i++ {
		a, _ := Square(i)
		sum += a
	}
	return sum
}
