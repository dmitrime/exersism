package secret

var vals = []struct {
	code   int
	action string
}{
	{1, "wink"},
	{2, "double blink"},
	{4, "close your eyes"},
	{8, "jump"},
}

func Handshake(n int) []string {
	if n <= 0 {
		return nil
	}
	res := []string{}
	for _, v := range vals {
		if n&v.code > 0 {
			res = append(res, v.action)
		}
	}
	if n&16 > 0 {
		for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
			res[i], res[j] = res[j], res[i]
		}
	}
	return res
}
