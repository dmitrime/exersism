package scrabble_score

import "strings"

func Score(t string) int {
	s := strings.ToUpper(t)
	sc := 0
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			if strings.ContainsRune("DG", rune(s[i])) {
				sc += 2
			} else if strings.ContainsRune("BCMP", rune(s[i])) {
				sc += 3
			} else if strings.ContainsRune("FHVWY", rune(s[i])) {
				sc += 4
			} else if s[i] == 'K' {
				sc += 5
			} else if strings.ContainsRune("JX", rune(s[i])) {
				sc += 8
			} else if strings.ContainsRune("QZ", rune(s[i])) {
				sc += 10
			} else {
				sc += 1
			}
		}
	}
	return sc
}
