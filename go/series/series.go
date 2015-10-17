package slice

func All(n int, s string) []string {
	var res []string
	if len(s) < n {
		return nil
	}
	for i := 0; i <= len(s)-n; i++ {
		res = append(res, s[i:i+n])
	}
	return res
}

func Frist(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return s, false
	}
	return s[:n], true
}
