package hamming

import "errors"

const TestVersion = 2

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("different length")
	}
	d := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			d++
		}
	}
	return d, nil
}
