package queenattack

import "errors"
import "math"

func CanQueenAttack(w string, b string) (bool, error) {
	if w != b && len(w) == 2 && len(b) == 2 &&
		w[0] >= 'a' && w[0] <= 'h' &&
		b[0] >= 'a' && b[0] <= 'h' &&
		w[1] >= '1' && w[1] <= '8' &&
		b[1] >= '1' && b[1] <= '8' {

		c1 := int(w[0] - 'a')
		c2 := int(b[0] - 'a')
		r1 := int(w[1] - '1')
		r2 := int(b[1] - '1')

		if c1 == c2 || r1 == r2 || math.Abs(float64(c1-c2)) == math.Abs(float64(r1-r2)) {
			return true, nil
		}

		return false, nil
	}
	return false, errors.New("")
}
