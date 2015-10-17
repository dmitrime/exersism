package etl

import "strings"

func Transform(in given) expectation {
	m := expectation{}
	for k, va := range in {
		for _, v := range va {
			m[strings.ToLower(v)] = k
		}
	}
	return m
}
