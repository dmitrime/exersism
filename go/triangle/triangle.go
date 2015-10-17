package triangle

type Kind string

const (
	Equ Kind = "equ"
	Iso      = "iso"
	Sca      = "sca"
	NaT      = "nat"
)

func KindFromSides(a float64, b float64, c float64) Kind {
	if a+b > c && a+c > b && b+c > a {
		if a == b && b == c && a > 0 {
			return Equ
		}
		if a == b && a != c || a == c && a != b || b == c && b != a {
			return Iso
		}
		return Sca
	}
	return NaT
}
